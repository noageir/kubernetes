/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"strings"
	"time"

	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/tools"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/golang/glog"
)

// RESTStorage is a generic interface for RESTful storage services
type RESTStorage interface {
	List(labels.Selector) (interface{}, error)
	Get(id string) (interface{}, error)
	Delete(id string) (<-chan interface{}, error)
	Extract(body []byte) (interface{}, error)
	Create(interface{}) (<-chan interface{}, error)
	Update(interface{}) (<-chan interface{}, error)
}

// WorkFunc is used to perform any time consuming work for an api call, after
// the input has been validated. Pass one of these to MakeAsync to create an
// appropriate return value for the Update, Delete, and Create methods.
type WorkFunc func() (result interface{}, err error)

// MakeAsync takes a function and executes it, delivering the result in the way required
// by RESTStorage's Update, Delete, and Create methods.
func MakeAsync(fn WorkFunc) <-chan interface{} {
	channel := make(chan interface{})
	go func() {
		defer util.HandleCrash()
		obj, err := fn()
		if err != nil {
			status := http.StatusInternalServerError
			switch {
			case tools.IsEtcdConflict(err):
				status = http.StatusConflict
			}
			channel <- &api.Status{
				Status:  api.StatusFailure,
				Details: err.Error(),
				Code:    status,
			}
		} else {
			channel <- obj
		}
		// 'close' is used to signal that no further values will
		// be written to the channel. Not strictly necessary, but
		// also won't hurt.
		close(channel)
	}()
	return channel
}

// APIServer is an HTTPHandler that delegates to RESTStorage objects.
// It handles URLs of the form:
// ${prefix}/${storage_key}[/${object_name}]
// Where 'prefix' is an arbitrary string, and 'storage_key' points to a RESTStorage object stored in storage.
//
// TODO: consider migrating this to go-restful which is a more full-featured version of the same thing.
type APIServer struct {
	prefix    string
	storage   map[string]RESTStorage
	ops       *Operations
	logserver http.Handler
}

// New creates a new APIServer object.
// 'storage' contains a map of handlers.
// 'prefix' is the hosting path prefix.
func New(storage map[string]RESTStorage, prefix string) *APIServer {
	return &APIServer{
		storage:   storage,
		prefix:    prefix,
		ops:       NewOperations(),
		logserver: http.StripPrefix("/logs/", http.FileServer(http.Dir("/var/log/"))),
	}
}

func (server *APIServer) handleIndex(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	// TODO: serve this out of a file?
	data := "<html><body>Welcome to Kubernetes</body></html>"
	fmt.Fprint(w, data)
}

func (server *APIServer) handleMinionReq(rawQuery string, w http.ResponseWriter) {
	// Expect rawQuery as: id=${minion}&query=/stats/<podid>/<containerName> or
	// id=${minion}&query=logs/
	values, err := url.ParseQuery(rawQuery)
	if err != nil {
		glog.Errorf("Invalid URL query: %s", rawQuery)
		return
	}

	minionHost := values.Get("id")
	queryUrl := "http://" + minionHost
	if strings.LastIndex(queryUrl, ":10250") < 0 {
		// No port information
		queryUrl += ":10250"
	}
	remote, err := url.Parse(queryUrl)
	if err != nil {
		glog.Errorf("Failed to parse %p as url: %s", queryUrl, err)
	}

	query := values.Get("query")
	newReq, err := http.NewRequest("GET", query, nil)
	if err != nil {
		glog.Errorf("Failed to create request: %s", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Transport = &minionTransport{}

	proxy.ServeHTTP(w, newReq)
}

type minionTransport struct{}

func (t *minionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(req)

	if !strings.HasPrefix(req.URL.Path, "/logs/") {
		// Do nothing, simply pass through
		return resp, err
	}

	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		// copying the response body did not work
		return nil, err
	}

	bodyNode := &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	}
	nodes, err := html.ParseFragment(bytes.NewBuffer(body), bodyNode)
	if err != nil {
		glog.Errorf("Failed to found <body> node: %v", err)
		return resp, err
	}

	// Define the method to traverse the doc tree and update href node to
	// point to correct minion
	var updateHRef func(*html.Node)
	updateHRef = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, attr := range n.Attr {
				if attr.Key == "href" {
					n.Attr[i].Val = "minion?id=" + req.URL.Host + "&query=" + req.URL.Path + attr.Val
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			updateHRef(c)
		}
	}

	newContent := &bytes.Buffer{}
	for _, n := range nodes {
		updateHRef(n)
		err = html.Render(newContent, n)
		if err != nil {
			glog.Errorf("Failed to render: %v", err)
		}
	}

	resp.Body = ioutil.NopCloser(newContent)
	// Update header node with new content-length
	resp.Header.Del("Content-Length")
	resp.ContentLength = int64(newContent.Len())

	return resp, err
}

// HTTP Handler interface
func (server *APIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "apiserver panic. Look in log for details.")
			glog.Infof("APIServer panic'd on %v %v: %#v\n%s\n", req.Method, req.RequestURI, x, debug.Stack())
		}
	}()
	defer MakeLogged(req, &w).StacktraceWhen(
		StatusIsNot(
			http.StatusOK,
			http.StatusAccepted,
			http.StatusConflict,
		),
	).Log()
	url, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		server.error(err, w)
		return
	}
	if url.Path == "/index.html" || url.Path == "/" || url.Path == "" {
		server.handleIndex(w)
		return
	}
	if url.Path == "/minion" {
		server.handleMinionReq(url.RawQuery, w)
		return
	}
	if url.Path == "/redirect" {
		req.URL.Path = "/redirect1"
		return
	}
	if strings.HasPrefix(url.Path, "/logs/") {
		server.logserver.ServeHTTP(w, req)
		return
	}
	if !strings.HasPrefix(url.Path, server.prefix) {
		server.notFound(req, w)
		return
	}
	requestParts := strings.Split(url.Path[len(server.prefix):], "/")[1:]
	if len(requestParts) < 1 {
		server.notFound(req, w)
		return
	}
	if requestParts[0] == "operations" {
		server.handleOperationRequest(requestParts[1:], w, req)
		return
	}
	storage := server.storage[requestParts[0]]
	if storage == nil {
		LogOf(w).Addf("'%v' has no storage object", requestParts[0])
		server.notFound(req, w)
		return
	}

	server.handleREST(requestParts, url, req, w, storage)
}

func (server *APIServer) notFound(req *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found: %#v", req)
}

func (server *APIServer) write(statusCode int, object interface{}, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	output, err := api.Encode(object)
	if err != nil {
		server.error(err, w)
		return
	}
	w.Write(output)
}

func (server *APIServer) error(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "Internal Error: %#v", err)
}

func (server *APIServer) readBody(req *http.Request) ([]byte, error) {
	defer req.Body.Close()
	return ioutil.ReadAll(req.Body)
}

// finishReq finishes up a request, waiting until the operation finishes or, after a timeout, creating an
// Operation to recieve the result and returning its ID down the writer.
func (server *APIServer) finishReq(out <-chan interface{}, sync bool, timeout time.Duration, w http.ResponseWriter) {
	op := server.ops.NewOperation(out)
	if sync {
		op.WaitFor(timeout)
	}
	obj, complete := op.StatusOrResult()
	if complete {
		status := http.StatusOK
		switch stat := obj.(type) {
		case api.Status:
			LogOf(w).Addf("programmer error: use *api.Status as a result, not api.Status.")
			if stat.Code != 0 {
				status = stat.Code
			}
		case *api.Status:
			if stat.Code != 0 {
				status = stat.Code
			}
		}
		server.write(status, obj, w)
	} else {
		server.write(http.StatusAccepted, obj, w)
	}
}

func parseTimeout(str string) time.Duration {
	if str != "" {
		timeout, err := time.ParseDuration(str)
		if err == nil {
			return timeout
		}
		glog.Errorf("Failed to parse: %#v '%s'", err, str)
	}
	return 30 * time.Second
}

// handleREST is the main dispatcher for the server.  It switches on the HTTP method, and then
// on path length, according to the following table:
//   Method     Path          Action
//   GET        /foo          list
//   GET        /foo/bar      get 'bar'
//   POST       /foo          create
//   PUT        /foo/bar      update 'bar'
//   DELETE     /foo/bar      delete 'bar'
// Returns 404 if the method/pattern doesn't match one of these entries
// The server accepts several query parameters:
//    sync=[false|true] Synchronous request (only applies to create, update, delete operations)
//    timeout=<duration> Timeout for synchronous requests, only applies if sync=true
//    labels=<label-selector> Used for filtering list operations
func (server *APIServer) handleREST(parts []string, requestURL *url.URL, req *http.Request, w http.ResponseWriter, storage RESTStorage) {
	sync := requestURL.Query().Get("sync") == "true"
	timeout := parseTimeout(requestURL.Query().Get("timeout"))
	switch req.Method {
	case "GET":
		switch len(parts) {
		case 1:
			selector, err := labels.ParseSelector(requestURL.Query().Get("labels"))
			if err != nil {
				server.error(err, w)
				return
			}
			list, err := storage.List(selector)
			if err != nil {
				server.error(err, w)
				return
			}
			server.write(http.StatusOK, list, w)
		case 2:
			item, err := storage.Get(parts[1])
			if err != nil {
				server.error(err, w)
				return
			}
			if item == nil {
				server.notFound(req, w)
				return
			}
			server.write(http.StatusOK, item, w)
		default:
			server.notFound(req, w)
		}
	case "POST":
		if len(parts) != 1 {
			server.notFound(req, w)
			return
		}
		body, err := server.readBody(req)
		if err != nil {
			server.error(err, w)
			return
		}
		obj, err := storage.Extract(body)
		if err != nil {
			server.error(err, w)
			return
		}
		out, err := storage.Create(obj)
		if err != nil {
			server.error(err, w)
			return
		}
		server.finishReq(out, sync, timeout, w)
	case "DELETE":
		if len(parts) != 2 {
			server.notFound(req, w)
			return
		}
		out, err := storage.Delete(parts[1])
		if err != nil {
			server.error(err, w)
			return
		}
		server.finishReq(out, sync, timeout, w)
	case "PUT":
		if len(parts) != 2 {
			server.notFound(req, w)
			return
		}
		body, err := server.readBody(req)
		if err != nil {
			server.error(err, w)
		}
		obj, err := storage.Extract(body)
		if err != nil {
			server.error(err, w)
			return
		}
		out, err := storage.Update(obj)
		if err != nil {
			server.error(err, w)
			return
		}
		server.finishReq(out, sync, timeout, w)
	default:
		server.notFound(req, w)
	}
}

func (server *APIServer) handleOperationRequest(parts []string, w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		server.notFound(req, w)
	}
	if len(parts) == 0 {
		// List outstanding operations.
		list := server.ops.List()
		server.write(http.StatusOK, list, w)
		return
	}

	op := server.ops.Get(parts[0])
	if op == nil {
		server.notFound(req, w)
	}

	obj, complete := op.StatusOrResult()
	if complete {
		server.write(http.StatusOK, obj, w)
	} else {
		server.write(http.StatusAccepted, obj, w)
	}
}

/*
Copyright 2018 The Kubernetes Authors.

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

package ssh

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
	"k8s.io/apimachinery/pkg/util/wait"
	sshutil "k8s.io/kubernetes/pkg/ssh"
)

// Result holds the execution result of SSH command
type Result struct {
	User   string
	Host   string
	Cmd    string
	Stdout string
	Stderr string
	Code   int
}

// GetSigner returns an ssh.Signer for the provider ("gce", etc.) that can be
// used to SSH to their nodes.
func GetSigner(provider string) (ssh.Signer, error) {
	// honor a consistent SSH key across all providers
	if path := os.Getenv("KUBE_SSH_KEY_PATH"); len(path) > 0 {
		return sshutil.MakePrivateKeySignerFromFile(path)
	}

	// Select the key itself to use. When implementing more providers here,
	// please also add them to any SSH tests that are disabled because of signer
	// support.
	keyfile := ""
	switch provider {
	case "gce", "gke", "kubemark":
		keyfile = os.Getenv("GCE_SSH_KEY")
		if keyfile == "" {
			keyfile = "google_compute_engine"
		}
	case "aws", "eks":
		keyfile = os.Getenv("AWS_SSH_KEY")
		if keyfile == "" {
			keyfile = "kube_aws_rsa"
		}
	case "local", "vsphere":
		keyfile = os.Getenv("LOCAL_SSH_KEY")
		if keyfile == "" {
			keyfile = "id_rsa"
		}
	case "skeleton":
		keyfile = os.Getenv("KUBE_SSH_KEY")
		if keyfile == "" {
			keyfile = "id_rsa"
		}
	default:
		return nil, fmt.Errorf("GetSigner(...) not implemented for %s", provider)
	}

	// Respect absolute paths for keys given by user, fallback to assuming
	// relative paths are in ~/.ssh
	if !filepath.IsAbs(keyfile) {
		keydir := filepath.Join(os.Getenv("HOME"), ".ssh")
		keyfile = filepath.Join(keydir, keyfile)
	}

	return sshutil.MakePrivateKeySignerFromFile(keyfile)
}

// SSH synchronously SSHs to a node running on provider and runs cmd. If there
// is no error performing the SSH, the stdout, stderr, and exit code are
// returned.
func SSH(cmd, host, provider string) (Result, error) {
	result := Result{Host: host, Cmd: cmd}

	// Get a signer for the provider.
	signer, err := GetSigner(provider)
	if err != nil {
		return result, fmt.Errorf("error getting signer for provider %s: '%v'", provider, err)
	}

	// RunSSHCommand will default to Getenv("USER") if user == "", but we're
	// defaulting here as well for logging clarity.
	result.User = os.Getenv("KUBE_SSH_USER")
	if result.User == "" {
		result.User = os.Getenv("USER")
	}

	if bastion := os.Getenv("KUBE_SSH_BASTION"); len(bastion) > 0 {
		stdout, stderr, code, err := runSSHCommandViaBastion(cmd, result.User, bastion, host, signer)
		result.Stdout = stdout
		result.Stderr = stderr
		result.Code = code
		return result, err
	}

	stdout, stderr, code, err := sshutil.RunSSHCommand(cmd, result.User, host, signer)
	result.Stdout = stdout
	result.Stderr = stderr
	result.Code = code

	return result, err
}

// runSSHCommandViaBastion returns the stdout, stderr, and exit code from running cmd on
// host as specific user, along with any SSH-level error. It uses an SSH proxy to connect
// to bastion, then via that tunnel connects to the remote host. Similar to
// sshutil.RunSSHCommand but scoped to the needs of the test infrastructure.
func runSSHCommandViaBastion(cmd, user, bastion, host string, signer ssh.Signer) (string, string, int, error) {
	// Setup the config, dial the server, and open a session.
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         150 * time.Second,
	}
	bastionClient, err := ssh.Dial("tcp", bastion, config)
	if err != nil {
		err = wait.Poll(5*time.Second, 20*time.Second, func() (bool, error) {
			fmt.Printf("error dialing %s@%s: '%v', retrying\n", user, bastion, err)
			if bastionClient, err = ssh.Dial("tcp", bastion, config); err != nil {
				return false, err
			}
			return true, nil
		})
	}
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting SSH client to %s@%s: %v", user, bastion, err)
	}
	defer bastionClient.Close()

	conn, err := bastionClient.Dial("tcp", host)
	if err != nil {
		return "", "", 0, fmt.Errorf("error dialing %s from bastion: %v", host, err)
	}
	defer conn.Close()

	ncc, chans, reqs, err := ssh.NewClientConn(conn, host, config)
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating forwarding connection %s from bastion: %v", host, err)
	}
	client := ssh.NewClient(ncc, chans, reqs)
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating session to %s@%s from bastion: '%v'", user, host, err)
	}
	defer session.Close()

	// Run the command.
	code := 0
	var bout, berr bytes.Buffer
	session.Stdout, session.Stderr = &bout, &berr
	if err = session.Run(cmd); err != nil {
		// Check whether the command failed to run or didn't complete.
		if exiterr, ok := err.(*ssh.ExitError); ok {
			// If we got an ExitError and the exit code is nonzero, we'll
			// consider the SSH itself successful (just that the command run
			// errored on the host).
			if code = exiterr.ExitStatus(); code != 0 {
				err = nil
			}
		} else {
			// Some other kind of error happened (e.g. an IOError); consider the
			// SSH unsuccessful.
			err = fmt.Errorf("failed running `%s` on %s@%s: '%v'", cmd, user, host, err)
		}
	}
	return bout.String(), berr.String(), code, err
}

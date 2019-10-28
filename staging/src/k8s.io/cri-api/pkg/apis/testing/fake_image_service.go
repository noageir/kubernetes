/*
Copyright 2016 The Kubernetes Authors.

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

package testing

import (
	"path"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type FakeImageService struct {
	sync.Mutex

	FakeImageSize uint64
	Called        []string
	Images        map[string]*runtimeapi.Image
	SandBoxes     map[string]string

	pulledImages []*pulledImage

	FakeFilesystemUsage []*runtimeapi.FilesystemUsage
}

func (r *FakeImageService) SetFakeImages(images []string) {
	r.Lock()
	defer r.Unlock()

	r.Images = make(map[string]*runtimeapi.Image)
	for _, image := range images {
		r.Images[image] = r.makeFakeImage(image)
	}
}

func (r *FakeImageService) SetFakeImageSize(size uint64) {
	r.Lock()
	defer r.Unlock()

	r.FakeImageSize = size
}

func (r *FakeImageService) SetFakeFilesystemUsage(usage []*runtimeapi.FilesystemUsage) {
	r.Lock()
	defer r.Unlock()

	r.FakeFilesystemUsage = usage
}

func NewFakeImageService() *FakeImageService {
	return &FakeImageService{
		Called: make([]string, 0),
		Images: make(map[string]*runtimeapi.Image),
	}
}

func (r *FakeImageService) makeFakeImage(image string) *runtimeapi.Image {
	return &runtimeapi.Image{
		Id:       image,
		Size_:    r.FakeImageSize,
		RepoTags: []string{image},
	}
}

// stringInSlice returns true if s is in list
func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
}

func (r *FakeImageService) ListImages(filter *runtimeapi.ImageFilter) ([]*runtimeapi.Image, error) {
	r.Lock()
	defer r.Unlock()

	r.Called = append(r.Called, "ListImages")

	images := make([]*runtimeapi.Image, 0)
	for _, img := range r.Images {
		if filter != nil && filter.Image != nil {
			if !stringInSlice(filter.Image.Image, img.RepoTags) {
				continue
			}
		}

		images = append(images, img)
	}
	return images, nil
}

func (r *FakeImageService) ImageStatus(image *runtimeapi.ImageSpec) (*runtimeapi.Image, error) {
	r.Lock()
	defer r.Unlock()

	r.Called = append(r.Called, "ImageStatus")

	return r.Images[image.Image], nil
}

func (r *FakeImageService) PullImage(image *runtimeapi.ImageSpec, auth *runtimeapi.AuthConfig, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	r.Lock()
	defer r.Unlock()

	r.Called = append(r.Called, "PullImage")

	if r.SandBoxes == nil {
		r.SandBoxes = make(map[string]string, 0)
	}

	r.SandBoxes[path.Join(podSandboxConfig.Metadata.Namespace, podSandboxConfig.Metadata.Name)] = podSandboxConfig.RuntimeHandler
	r.pulledImages = append(r.pulledImages, &pulledImage{imageSpec: image, authConfig: auth})
	// ImageID should be randomized for real container runtime, but here just use
	// image's name for easily making fake images.
	imageID := image.Image
	if _, ok := r.Images[imageID]; !ok {
		r.Images[imageID] = r.makeFakeImage(image.Image)
	}

	return imageID, nil
}

func (r *FakeImageService) ListSandBoxes() map[string]string {
	return r.SandBoxes
}

func (r *FakeImageService) RemoveImage(image *runtimeapi.ImageSpec) error {
	r.Lock()
	defer r.Unlock()

	r.Called = append(r.Called, "RemoveImage")

	// Remove the image
	delete(r.Images, image.Image)

	return nil
}

// ImageFsInfo returns information of the filesystem that is used to store images.
func (r *FakeImageService) ImageFsInfo() ([]*runtimeapi.FilesystemUsage, error) {
	r.Lock()
	defer r.Unlock()

	r.Called = append(r.Called, "ImageFsInfo")

	return r.FakeFilesystemUsage, nil
}

func (r *FakeImageService) AssertImagePulledWithAuth(t *testing.T, image *runtimeapi.ImageSpec, auth *runtimeapi.AuthConfig, failMsg string) {
	r.Lock()
	defer r.Unlock()
	expected := &pulledImage{imageSpec: image, authConfig: auth}
	assert.Contains(t, r.pulledImages, expected, failMsg)
}

type pulledImage struct {
	imageSpec  *runtimeapi.ImageSpec
	authConfig *runtimeapi.AuthConfig
}

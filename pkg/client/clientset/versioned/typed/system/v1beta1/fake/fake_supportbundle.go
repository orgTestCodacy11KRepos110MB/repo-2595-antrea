// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "antrea.io/antrea/pkg/apis/system/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeSupportBundles implements SupportBundleInterface
type FakeSupportBundles struct {
	Fake *FakeSystemV1beta1
}

var supportbundlesResource = schema.GroupVersionResource{Group: "system.antrea.io", Version: "v1beta1", Resource: "supportbundles"}

var supportbundlesKind = schema.GroupVersionKind{Group: "system.antrea.io", Version: "v1beta1", Kind: "SupportBundle"}

// Get takes name of the supportBundle, and returns the corresponding supportBundle object, and an error if there is any.
func (c *FakeSupportBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(supportbundlesResource, name), &v1beta1.SupportBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

// Create takes the representation of a supportBundle and creates it.  Returns the server's representation of the supportBundle, and an error, if there is any.
func (c *FakeSupportBundles) Create(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.CreateOptions) (result *v1beta1.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(supportbundlesResource, supportBundle), &v1beta1.SupportBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

// Delete takes name of the supportBundle and deletes it. Returns an error if one occurs.
func (c *FakeSupportBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(supportbundlesResource, name, opts), &v1beta1.SupportBundle{})
	return err
}

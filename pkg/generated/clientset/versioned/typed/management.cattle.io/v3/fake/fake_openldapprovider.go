/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenLdapProviders implements OpenLdapProviderInterface
type FakeOpenLdapProviders struct {
	Fake *FakeManagementV3
}

var openldapprovidersResource = v3.SchemeGroupVersion.WithResource("openldapproviders")

var openldapprovidersKind = v3.SchemeGroupVersion.WithKind("OpenLdapProvider")

// Get takes name of the openLdapProvider, and returns the corresponding openLdapProvider object, and an error if there is any.
func (c *FakeOpenLdapProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.OpenLdapProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(openldapprovidersResource, name), &v3.OpenLdapProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.OpenLdapProvider), err
}

// List takes label and field selectors, and returns the list of OpenLdapProviders that match those selectors.
func (c *FakeOpenLdapProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.OpenLdapProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(openldapprovidersResource, openldapprovidersKind, opts), &v3.OpenLdapProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.OpenLdapProviderList{ListMeta: obj.(*v3.OpenLdapProviderList).ListMeta}
	for _, item := range obj.(*v3.OpenLdapProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openLdapProviders.
func (c *FakeOpenLdapProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(openldapprovidersResource, opts))
}

// Create takes the representation of a openLdapProvider and creates it.  Returns the server's representation of the openLdapProvider, and an error, if there is any.
func (c *FakeOpenLdapProviders) Create(ctx context.Context, openLdapProvider *v3.OpenLdapProvider, opts v1.CreateOptions) (result *v3.OpenLdapProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(openldapprovidersResource, openLdapProvider), &v3.OpenLdapProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.OpenLdapProvider), err
}

// Update takes the representation of a openLdapProvider and updates it. Returns the server's representation of the openLdapProvider, and an error, if there is any.
func (c *FakeOpenLdapProviders) Update(ctx context.Context, openLdapProvider *v3.OpenLdapProvider, opts v1.UpdateOptions) (result *v3.OpenLdapProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(openldapprovidersResource, openLdapProvider), &v3.OpenLdapProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.OpenLdapProvider), err
}

// Delete takes name of the openLdapProvider and deletes it. Returns an error if one occurs.
func (c *FakeOpenLdapProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(openldapprovidersResource, name, opts), &v3.OpenLdapProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenLdapProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(openldapprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.OpenLdapProviderList{})
	return err
}

// Patch applies the patch and returns the patched openLdapProvider.
func (c *FakeOpenLdapProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.OpenLdapProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openldapprovidersResource, name, pt, data, subresources...), &v3.OpenLdapProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.OpenLdapProvider), err
}
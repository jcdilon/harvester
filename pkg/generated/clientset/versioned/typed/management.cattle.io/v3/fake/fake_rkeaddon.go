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

// FakeRkeAddons implements RkeAddonInterface
type FakeRkeAddons struct {
	Fake *FakeManagementV3
	ns   string
}

var rkeaddonsResource = v3.SchemeGroupVersion.WithResource("rkeaddons")

var rkeaddonsKind = v3.SchemeGroupVersion.WithKind("RkeAddon")

// Get takes name of the rkeAddon, and returns the corresponding rkeAddon object, and an error if there is any.
func (c *FakeRkeAddons) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.RkeAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rkeaddonsResource, c.ns, name), &v3.RkeAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeAddon), err
}

// List takes label and field selectors, and returns the list of RkeAddons that match those selectors.
func (c *FakeRkeAddons) List(ctx context.Context, opts v1.ListOptions) (result *v3.RkeAddonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rkeaddonsResource, rkeaddonsKind, c.ns, opts), &v3.RkeAddonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.RkeAddonList{ListMeta: obj.(*v3.RkeAddonList).ListMeta}
	for _, item := range obj.(*v3.RkeAddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rkeAddons.
func (c *FakeRkeAddons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rkeaddonsResource, c.ns, opts))

}

// Create takes the representation of a rkeAddon and creates it.  Returns the server's representation of the rkeAddon, and an error, if there is any.
func (c *FakeRkeAddons) Create(ctx context.Context, rkeAddon *v3.RkeAddon, opts v1.CreateOptions) (result *v3.RkeAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rkeaddonsResource, c.ns, rkeAddon), &v3.RkeAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeAddon), err
}

// Update takes the representation of a rkeAddon and updates it. Returns the server's representation of the rkeAddon, and an error, if there is any.
func (c *FakeRkeAddons) Update(ctx context.Context, rkeAddon *v3.RkeAddon, opts v1.UpdateOptions) (result *v3.RkeAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rkeaddonsResource, c.ns, rkeAddon), &v3.RkeAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeAddon), err
}

// Delete takes name of the rkeAddon and deletes it. Returns an error if one occurs.
func (c *FakeRkeAddons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rkeaddonsResource, c.ns, name, opts), &v3.RkeAddon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRkeAddons) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rkeaddonsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.RkeAddonList{})
	return err
}

// Patch applies the patch and returns the patched rkeAddon.
func (c *FakeRkeAddons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RkeAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rkeaddonsResource, c.ns, name, pt, data, subresources...), &v3.RkeAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeAddon), err
}
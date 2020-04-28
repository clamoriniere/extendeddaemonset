// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/datadog/extendeddaemonset/pkg/apis/datadoghq/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExtendedNodes implements ExtendedNodeInterface
type FakeExtendedNodes struct {
	Fake *FakeDatadoghqV1alpha1
	ns   string
}

var extendednodesResource = schema.GroupVersionResource{Group: "datadoghq.com", Version: "v1alpha1", Resource: "extendednodes"}

var extendednodesKind = schema.GroupVersionKind{Group: "datadoghq.com", Version: "v1alpha1", Kind: "ExtendedNode"}

// Get takes name of the extendedNode, and returns the corresponding extendedNode object, and an error if there is any.
func (c *FakeExtendedNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.ExtendedNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(extendednodesResource, c.ns, name), &v1alpha1.ExtendedNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExtendedNode), err
}

// List takes label and field selectors, and returns the list of ExtendedNodes that match those selectors.
func (c *FakeExtendedNodes) List(opts v1.ListOptions) (result *v1alpha1.ExtendedNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(extendednodesResource, extendednodesKind, c.ns, opts), &v1alpha1.ExtendedNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ExtendedNodeList{ListMeta: obj.(*v1alpha1.ExtendedNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ExtendedNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested extendedNodes.
func (c *FakeExtendedNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(extendednodesResource, c.ns, opts))

}

// Create takes the representation of a extendedNode and creates it.  Returns the server's representation of the extendedNode, and an error, if there is any.
func (c *FakeExtendedNodes) Create(extendedNode *v1alpha1.ExtendedNode) (result *v1alpha1.ExtendedNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(extendednodesResource, c.ns, extendedNode), &v1alpha1.ExtendedNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExtendedNode), err
}

// Update takes the representation of a extendedNode and updates it. Returns the server's representation of the extendedNode, and an error, if there is any.
func (c *FakeExtendedNodes) Update(extendedNode *v1alpha1.ExtendedNode) (result *v1alpha1.ExtendedNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(extendednodesResource, c.ns, extendedNode), &v1alpha1.ExtendedNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExtendedNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExtendedNodes) UpdateStatus(extendedNode *v1alpha1.ExtendedNode) (*v1alpha1.ExtendedNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(extendednodesResource, "status", c.ns, extendedNode), &v1alpha1.ExtendedNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExtendedNode), err
}

// Delete takes name of the extendedNode and deletes it. Returns an error if one occurs.
func (c *FakeExtendedNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(extendednodesResource, c.ns, name), &v1alpha1.ExtendedNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExtendedNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(extendednodesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ExtendedNodeList{})
	return err
}

// Patch applies the patch and returns the patched extendedNode.
func (c *FakeExtendedNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ExtendedNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(extendednodesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ExtendedNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExtendedNode), err
}

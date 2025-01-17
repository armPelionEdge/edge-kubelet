/*
Copyright 2018-2020, Arm Limited and affiliates.

Licensed under the Apache License, Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an AS IS BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// NodesGetter has a method to return a NodeInterface.
// A group's client should implement this interface.
type NodesGetter interface {
	Nodes(aid string) NodeInterface
}

// NodeInterface has methods to work with Node resources.
type NodeInterface interface {
	Create(*v1.Node) (*v1.Node, error)
	Update(*v1.Node) (*v1.Node, error)
	UpdateStatus(*v1.Node) (*v1.Node, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Node, error)
	List(opts meta_v1.ListOptions) (*v1.NodeList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Node, err error)
	NodeExpansion
}

// nodes implements NodeInterface
type nodes struct {
	client rest.Interface
	aid    string
}

// newNodes returns a Nodes
func newNodes(c *ArgusV1Client, aid string) *nodes {
	return &nodes{
		client: c.RESTClient(),
		aid:    aid,
	}
}

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *nodes) Get(name string, options meta_v1.GetOptions) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Get().
		AccountID(c.aid).
		Resource("nodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *nodes) List(opts meta_v1.ListOptions) (result *v1.NodeList, err error) {
	result = &v1.NodeList{}
	err = c.client.Get().
		AccountID(c.aid).
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *nodes) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		AccountID(c.aid).
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *nodes) Create(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Post().
		AccountID(c.aid).
		Resource("nodes").
		Body(node).
		Do().
		Into(result)
	return
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *nodes) Update(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Put().
		AccountID(c.aid).
		Resource("nodes").
		Name(node.Name).
		Body(node).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodes) UpdateStatus(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Put().
		AccountID(c.aid).
		Resource("nodes").
		Name(node.Name).
		SubResource("status").
		Body(node).
		Do().
		Into(result)
	return
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *nodes) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		AccountID(c.aid).
		Resource("nodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodes) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		AccountID(c.aid).
		Resource("nodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched node.
func (c *nodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Patch(pt).
		AccountID(c.aid).
		Resource("nodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

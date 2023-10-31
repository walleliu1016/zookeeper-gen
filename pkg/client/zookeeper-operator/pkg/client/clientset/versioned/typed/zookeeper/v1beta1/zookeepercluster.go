/*
Copyright The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"
	v1beta1 "zookeeper-operator/pkg/apis/zookeeper/v1beta1"
	scheme "zookeeper-operator/pkg/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ZookeeperClustersGetter has a method to return a ZookeeperClusterInterface.
// A group's client should implement this interface.
type ZookeeperClustersGetter interface {
	ZookeeperClusters(namespace string) ZookeeperClusterInterface
}

// ZookeeperClusterInterface has methods to work with ZookeeperCluster resources.
type ZookeeperClusterInterface interface {
	Create(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.CreateOptions) (*v1beta1.ZookeeperCluster, error)
	Update(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.UpdateOptions) (*v1beta1.ZookeeperCluster, error)
	UpdateStatus(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.UpdateOptions) (*v1beta1.ZookeeperCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ZookeeperCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ZookeeperClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ZookeeperCluster, err error)
	ZookeeperClusterExpansion
}

// zookeeperClusters implements ZookeeperClusterInterface
type zookeeperClusters struct {
	client rest.Interface
	ns     string
}

// newZookeeperClusters returns a ZookeeperClusters
func newZookeeperClusters(c *ZookeeperV1beta1Client, namespace string) *zookeeperClusters {
	return &zookeeperClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the zookeeperCluster, and returns the corresponding zookeeperCluster object, and an error if there is any.
func (c *zookeeperClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ZookeeperCluster, err error) {
	result = &v1beta1.ZookeeperCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ZookeeperClusters that match those selectors.
func (c *zookeeperClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ZookeeperClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ZookeeperClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested zookeeperClusters.
func (c *zookeeperClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a zookeeperCluster and creates it.  Returns the server's representation of the zookeeperCluster, and an error, if there is any.
func (c *zookeeperClusters) Create(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.CreateOptions) (result *v1beta1.ZookeeperCluster, err error) {
	result = &v1beta1.ZookeeperCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(zookeeperCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a zookeeperCluster and updates it. Returns the server's representation of the zookeeperCluster, and an error, if there is any.
func (c *zookeeperClusters) Update(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.UpdateOptions) (result *v1beta1.ZookeeperCluster, err error) {
	result = &v1beta1.ZookeeperCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		Name(zookeeperCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(zookeeperCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *zookeeperClusters) UpdateStatus(ctx context.Context, zookeeperCluster *v1beta1.ZookeeperCluster, opts v1.UpdateOptions) (result *v1beta1.ZookeeperCluster, err error) {
	result = &v1beta1.ZookeeperCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		Name(zookeeperCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(zookeeperCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the zookeeperCluster and deletes it. Returns an error if one occurs.
func (c *zookeeperClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *zookeeperClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zookeeperclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched zookeeperCluster.
func (c *zookeeperClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ZookeeperCluster, err error) {
	result = &v1beta1.ZookeeperCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("zookeeperclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

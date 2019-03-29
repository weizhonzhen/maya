/*
Copyright 2019 The OpenEBS Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/upgrade/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/generated/openebs.io/upgrade/v1alpha1/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UpgradeResultsGetter has a method to return a UpgradeResultInterface.
// A group's client should implement this interface.
type UpgradeResultsGetter interface {
	UpgradeResults(namespace string) UpgradeResultInterface
}

// UpgradeResultInterface has methods to work with UpgradeResult resources.
type UpgradeResultInterface interface {
	Create(*v1alpha1.UpgradeResult) (*v1alpha1.UpgradeResult, error)
	Update(*v1alpha1.UpgradeResult) (*v1alpha1.UpgradeResult, error)
	UpdateStatus(*v1alpha1.UpgradeResult) (*v1alpha1.UpgradeResult, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UpgradeResult, error)
	List(opts v1.ListOptions) (*v1alpha1.UpgradeResultList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UpgradeResult, err error)
	UpgradeResultExpansion
}

// upgradeResults implements UpgradeResultInterface
type upgradeResults struct {
	client rest.Interface
	ns     string
}

// newUpgradeResults returns a UpgradeResults
func newUpgradeResults(c *OpenebsV1alpha1Client, namespace string) *upgradeResults {
	return &upgradeResults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the upgradeResult, and returns the corresponding upgradeResult object, and an error if there is any.
func (c *upgradeResults) Get(name string, options v1.GetOptions) (result *v1alpha1.UpgradeResult, err error) {
	result = &v1alpha1.UpgradeResult{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upgraderesults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UpgradeResults that match those selectors.
func (c *upgradeResults) List(opts v1.ListOptions) (result *v1alpha1.UpgradeResultList, err error) {
	result = &v1alpha1.UpgradeResultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upgraderesults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested upgradeResults.
func (c *upgradeResults) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("upgraderesults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a upgradeResult and creates it.  Returns the server's representation of the upgradeResult, and an error, if there is any.
func (c *upgradeResults) Create(upgradeResult *v1alpha1.UpgradeResult) (result *v1alpha1.UpgradeResult, err error) {
	result = &v1alpha1.UpgradeResult{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("upgraderesults").
		Body(upgradeResult).
		Do().
		Into(result)
	return
}

// Update takes the representation of a upgradeResult and updates it. Returns the server's representation of the upgradeResult, and an error, if there is any.
func (c *upgradeResults) Update(upgradeResult *v1alpha1.UpgradeResult) (result *v1alpha1.UpgradeResult, err error) {
	result = &v1alpha1.UpgradeResult{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upgraderesults").
		Name(upgradeResult.Name).
		Body(upgradeResult).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *upgradeResults) UpdateStatus(upgradeResult *v1alpha1.UpgradeResult) (result *v1alpha1.UpgradeResult, err error) {
	result = &v1alpha1.UpgradeResult{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upgraderesults").
		Name(upgradeResult.Name).
		SubResource("status").
		Body(upgradeResult).
		Do().
		Into(result)
	return
}

// Delete takes name of the upgradeResult and deletes it. Returns an error if one occurs.
func (c *upgradeResults) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upgraderesults").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *upgradeResults) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upgraderesults").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched upgradeResult.
func (c *upgradeResults) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UpgradeResult, err error) {
	result = &v1alpha1.UpgradeResult{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("upgraderesults").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
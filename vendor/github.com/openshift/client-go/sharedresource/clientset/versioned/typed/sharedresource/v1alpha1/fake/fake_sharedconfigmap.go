// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openshift/api/sharedresource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharedConfigMaps implements SharedConfigMapInterface
type FakeSharedConfigMaps struct {
	Fake *FakeSharedresourceV1alpha1
}

var sharedconfigmapsResource = schema.GroupVersionResource{Group: "sharedresource.openshift.io", Version: "v1alpha1", Resource: "sharedconfigmaps"}

var sharedconfigmapsKind = schema.GroupVersionKind{Group: "sharedresource.openshift.io", Version: "v1alpha1", Kind: "SharedConfigMap"}

// Get takes name of the sharedConfigMap, and returns the corresponding sharedConfigMap object, and an error if there is any.
func (c *FakeSharedConfigMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SharedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sharedconfigmapsResource, name), &v1alpha1.SharedConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfigMap), err
}

// List takes label and field selectors, and returns the list of SharedConfigMaps that match those selectors.
func (c *FakeSharedConfigMaps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SharedConfigMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sharedconfigmapsResource, sharedconfigmapsKind, opts), &v1alpha1.SharedConfigMapList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharedConfigMapList{ListMeta: obj.(*v1alpha1.SharedConfigMapList).ListMeta}
	for _, item := range obj.(*v1alpha1.SharedConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedConfigMaps.
func (c *FakeSharedConfigMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sharedconfigmapsResource, opts))
}

// Create takes the representation of a sharedConfigMap and creates it.  Returns the server's representation of the sharedConfigMap, and an error, if there is any.
func (c *FakeSharedConfigMaps) Create(ctx context.Context, sharedConfigMap *v1alpha1.SharedConfigMap, opts v1.CreateOptions) (result *v1alpha1.SharedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sharedconfigmapsResource, sharedConfigMap), &v1alpha1.SharedConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfigMap), err
}

// Update takes the representation of a sharedConfigMap and updates it. Returns the server's representation of the sharedConfigMap, and an error, if there is any.
func (c *FakeSharedConfigMaps) Update(ctx context.Context, sharedConfigMap *v1alpha1.SharedConfigMap, opts v1.UpdateOptions) (result *v1alpha1.SharedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sharedconfigmapsResource, sharedConfigMap), &v1alpha1.SharedConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfigMap), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedConfigMaps) UpdateStatus(ctx context.Context, sharedConfigMap *v1alpha1.SharedConfigMap, opts v1.UpdateOptions) (*v1alpha1.SharedConfigMap, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(sharedconfigmapsResource, "status", sharedConfigMap), &v1alpha1.SharedConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfigMap), err
}

// Delete takes name of the sharedConfigMap and deletes it. Returns an error if one occurs.
func (c *FakeSharedConfigMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(sharedconfigmapsResource, name), &v1alpha1.SharedConfigMap{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedConfigMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sharedconfigmapsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharedConfigMapList{})
	return err
}

// Patch applies the patch and returns the patched sharedConfigMap.
func (c *FakeSharedConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SharedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sharedconfigmapsResource, name, pt, data, subresources...), &v1alpha1.SharedConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfigMap), err
}
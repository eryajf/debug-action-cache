/*
Copyright 2025 Rancher Labs, Inc.

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

package v1

import (
	context "context"

	catalogcattleiov1 "github.com/rancher/rancher/pkg/apis/catalog.cattle.io/v1"
	scheme "github.com/rancher/rancher/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// OperationsGetter has a method to return a OperationInterface.
// A group's client should implement this interface.
type OperationsGetter interface {
	Operations(namespace string) OperationInterface
}

// OperationInterface has methods to work with Operation resources.
type OperationInterface interface {
	Create(ctx context.Context, operation *catalogcattleiov1.Operation, opts metav1.CreateOptions) (*catalogcattleiov1.Operation, error)
	Update(ctx context.Context, operation *catalogcattleiov1.Operation, opts metav1.UpdateOptions) (*catalogcattleiov1.Operation, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, operation *catalogcattleiov1.Operation, opts metav1.UpdateOptions) (*catalogcattleiov1.Operation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*catalogcattleiov1.Operation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*catalogcattleiov1.OperationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *catalogcattleiov1.Operation, err error)
	OperationExpansion
}

// operations implements OperationInterface
type operations struct {
	*gentype.ClientWithList[*catalogcattleiov1.Operation, *catalogcattleiov1.OperationList]
}

// newOperations returns a Operations
func newOperations(c *CatalogV1Client, namespace string) *operations {
	return &operations{
		gentype.NewClientWithList[*catalogcattleiov1.Operation, *catalogcattleiov1.OperationList](
			"operations",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *catalogcattleiov1.Operation { return &catalogcattleiov1.Operation{} },
			func() *catalogcattleiov1.OperationList { return &catalogcattleiov1.OperationList{} },
		),
	}
}

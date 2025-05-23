/*
Copyright The Karmada Authors.

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
	context "context"

	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClustersGetter has a method to return a ClusterInterface.
// A group's client should implement this interface.
type ClustersGetter interface {
	Clusters() ClusterInterface
}

// ClusterInterface has methods to work with Cluster resources.
type ClusterInterface interface {
	Create(ctx context.Context, cluster *clusterv1alpha1.Cluster, opts v1.CreateOptions) (*clusterv1alpha1.Cluster, error)
	Update(ctx context.Context, cluster *clusterv1alpha1.Cluster, opts v1.UpdateOptions) (*clusterv1alpha1.Cluster, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cluster *clusterv1alpha1.Cluster, opts v1.UpdateOptions) (*clusterv1alpha1.Cluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*clusterv1alpha1.Cluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*clusterv1alpha1.ClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusterv1alpha1.Cluster, err error)
	ClusterExpansion
}

// clusters implements ClusterInterface
type clusters struct {
	*gentype.ClientWithList[*clusterv1alpha1.Cluster, *clusterv1alpha1.ClusterList]
}

// newClusters returns a Clusters
func newClusters(c *ClusterV1alpha1Client) *clusters {
	return &clusters{
		gentype.NewClientWithList[*clusterv1alpha1.Cluster, *clusterv1alpha1.ClusterList](
			"clusters",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *clusterv1alpha1.Cluster { return &clusterv1alpha1.Cluster{} },
			func() *clusterv1alpha1.ClusterList { return &clusterv1alpha1.ClusterList{} },
		),
	}
}

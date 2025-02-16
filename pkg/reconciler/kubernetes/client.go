package kubernetes

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes"
)

type Resource struct {
	Kind      string
	Name      string
	Namespace string
}

func (r *Resource) String() string {
	return fmt.Sprintf("Resource [Kind:%s,Namespace:%s,Name:%s]", r.Kind, r.Namespace, r.Name)
}

type ResourceInterceptor interface {
	Intercept(resource *unstructured.Unstructured) error
}

type Client interface {
	Deploy(manifest string, interceptors ...ResourceInterceptor) ([]*Resource, error)
	Delete(manifest string) error
	Clientset() (*kubernetes.Clientset, error)
}

func NewKubernetesClient(kubeconfig string, debug bool) (Client, error) {
	return newKubeClientAdapter(kubeconfig, debug)
}

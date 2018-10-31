package v1

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/factory"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	APIVersion = types.APIVersion{
		Group:   "some.api.group",
		Version: "v1",
		Path:    "/v1-someapi",
	}
	Schemas = factory.
		Schemas(&APIVersion).
		MustImport(&APIVersion, Foo{})
)

type Foo struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FooSpec `json:"spec"`
}

type FooSpec struct {
	Option bool `json:"option"`
}

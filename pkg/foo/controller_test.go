package foo

import (
	"%PKG%/pkg/apis/some.api.group/v1"
	fooFakes "%PKG%/pkg/generated/controllers/some.api.group/v1/fakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFooControllerOnChange(t *testing.T) {
	assert := assert.New(t)
	controller := NewMockFooController()
	foo, key := NewFoo()
	returnedFoo, _ := controller.OnFooChange(key, foo)

	//write asserts
	assert.Equal("foo", returnedFoo.Name)
}

func TestFooControllerOnRemove(t *testing.T) {
	assert := assert.New(t)
	controller := NewMockFooController()
	foo, key := NewFoo()
	returnedFoo, _ := controller.OnFooRemove(key, foo)

	//write asserts
	assert.Equal("foo", returnedFoo.Name)
}

func NewFoo() (foo *v1.Foo, key string) {
	foo = v1.NewFoo("kube-system", "foo", v1.Foo{})
	key = foo.Namespace + "/" + foo.Name

	return
}

func NewMockFooController() Controller {
	foos := &fooFakes.FooControllerMock{
		UpdateFunc: func(in1 *v1.Foo) (*v1.Foo, error) {
			return in1, nil
		},
	}

	return Controller{
		foos: foos,
	}
}

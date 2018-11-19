package foo

import (
	"context"

	"%PKG%/types/apis/some.api.group/v1"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
)

func Register(ctx context.Context, client v1.Interface) error {
	fl := &fooLifecycle{
		fooClient: client.Foos(""),
		fooLister: client.Foos("").Controller().Lister(),
	}

	client.Foos("").AddHandler(ctx, "foo controller", SyncHandler)
	client.Foos("").AddLifecycle(ctx, "foo controller", fl)
	return nil
}

func SyncHandler(key string, obj *v1.Foo) (runtime.Object, error) {
	// Called anytime something changes, obj will be nil on delete
	logrus.Infof("Sync handler called %s %v", key, obj)
	return obj, nil
}

type fooLifecycle struct {
	fooClient v1.FooInterface
	fooLister v1.FooLister
}

func (f *fooLifecycle) Create(obj *v1.Foo) (runtime.Object, error) {
	logrus.Infof("Created: %v", obj)
	return obj, nil
}

func (f *fooLifecycle) Remove(obj *v1.Foo) (runtime.Object, error) {
	logrus.Infof("Finalizer: %v", obj)
	return obj, nil
}

func (f *fooLifecycle) Updated(obj *v1.Foo) (runtime.Object, error) {
	logrus.Infof("Updated: %v", obj)
	return obj, nil
}

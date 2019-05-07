//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go

package main

import (
	"context"
	"flag"
	"%PKG%/pkg/foo"
	"%PKG%/pkg/generated/controllers/some.api.group"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/rancher/wrangler/pkg/start"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	ctx := signals.SetupSignalHandler(context.Background())

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	foos, err := some.NewFactoryFromConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building sample controllers: %s", err.Error())
	}

	foo.Register(ctx, foos.Some().V1().Foo())

	if err := start.All(ctx, 2, foos); err != nil {
		klog.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

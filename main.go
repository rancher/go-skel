//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go

package main

import (
	"context"
	"flag"
	"%PKG%/pkg/foo"
	"%PKG%/pkg/generated/controllers/some.api.group"
	"github.com/rancher/wrangler/pkg/resolvehome"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/rancher/wrangler/pkg/start"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var (
	VERSION = "v0.0.0-dev"
)

func main() {
	app := cli.NewApp()
	app.Name = "testy"
	app.Version = VERSION
	app.Usage = "testy needs help!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kubeconfig",
			EnvVar: "KUBECONFIG",
			Value:  "${HOME}/.kube/config",
		},
		cli.StringFlag{
			Name:   "masterurl",
			EnvVar: "MASTERURL",
			Value:  "",
		},
	}
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) {
	flag.Parse()

	logrus.Info("Starting controller")
	ctx := signals.SetupSignalHandler(context.Background())

	kubeconfig, err := resolvehome.Resolve(c.String("kubeconfig"))
	if err != nil {
		logrus.Fatalf("Error resolving home dir: %s", err.Error())
	}
	masterurl := c.String("masterurl")

	cfg, err := clientcmd.BuildConfigFromFlags(masterurl, kubeconfig)
	if err != nil {
		logrus.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	foos, err := some.NewFactoryFromConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error building sample controllers: %s", err.Error())
	}

	foo.Register(ctx, foos.Some().V1().Foo())

	if err := start.All(ctx, 2, foos); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}

//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go

package main

import (
	"context"
	"flag"
	"%PKG%/pkg/foo"
	"%PKG%/pkg/generated/controllers/some.api.group"
	"%PKG%/pkg/version"
	"github.com/rancher/wrangler/pkg/kubeconfig"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/rancher/wrangler/pkg/start"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	KubeConfig string
)

func main() {
	app := cli.NewApp()
	app.Name = "testy"
	app.Version = version.FriendlyVersion()
	app.Usage = "testy needs help!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kubeconfig",
			EnvVar: "KUBECONFIG",
			Destination: &KubeConfig,
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

	kubeConfig, err := kubeconfig.GetNonInteractiveClientConfig(KubeConfig).ClientConfig()
	if err != nil {
		logrus.Fatalf("failed to find kubeconfig: %v", err)
	}

	foos, err := some.NewFactoryFromConfig(kubeConfig)
	if err != nil {
		logrus.Fatalf("Error building sample controllers: %s", err.Error())
	}

	foo.Register(ctx, foos.Some().V1().Foo())

	if err := start.All(ctx, 2, foos); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}

package main

import (
	"%PKG%/pkg/apis/some.api.group/v1"
	"github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "%PKG%/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"some.api.group": {
				Types: []interface{}{
					v1.Foo{},
				},
				GenerateTypes: true,
			},
		},
	})
}

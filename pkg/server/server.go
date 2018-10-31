package server

import (
	"context"

	"%PKG%/pkg/controllers/foo"
	"%PKG%/types/apis/some.api.group/v1"
	"github.com/rancher/norman"
	"github.com/rancher/norman/types"
)

func Config() *norman.Config {
	return &norman.Config{
		Name: "%APP%",
		Schemas: []*types.Schemas{
			v1.Schemas,
		},

		CRDs: map[*types.APIVersion][]string{
			&v1.APIVersion: {
				v1.FooGroupVersionKind.Kind,
			},
		},

		Clients: []norman.ClientFactory{
			v1.Factory,
		},

		MasterControllers: []norman.ControllerRegister{
			func(ctx context.Context) error {
				return foo.Register(ctx, v1.From(ctx))
			},
		},
	}
}

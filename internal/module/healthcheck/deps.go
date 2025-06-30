package healthcheck

import (
	"go-modular-sample/internal/module/healthcheck/present/api"
	"go-modular-sample/internal/shared/module"
)

func UseCases() module.UseCases {
	return module.UseCases{}
}

func Deps() module.Deps {
	return module.Deps{}
}

func Routes() module.Routes {
	return module.Routes{
		api.NewHealthCheckController,
	}
}

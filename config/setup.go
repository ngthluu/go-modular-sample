package config

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go-modular-sample/internal/shared/api"
	"go-modular-sample/internal/shared/db"

	"go.uber.org/fx"
)

type AppInstance struct{}

func setupLifeCycle(lc fx.Lifecycle, server *http.Server, db db.Db) *AppInstance {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", server.Addr)
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer db.GetConn().Close()
			return server.Shutdown(ctx)
		},
	})
	return &AppInstance{}
}

func setupInternalDeps() []any {
	var opts []any
	deps := GetAllInternalDeps()
	for inf, impl := range deps {
		opts = append(opts, fx.Annotate(
			impl,
			fx.As(inf),
		))
	}
	return opts
}

func setupModuleConfigs() []any {
	var opts []any
	configs := GetAllModuleConfigs()
	for _, cog := range configs {
		// Setup usecases
		opts = append(opts, cog.UseCases...)

		// Setup deps
		for inf, impl := range cog.Deps {
			opts = append(opts, fx.Annotate(
				impl,
				fx.As(inf),
			))
		}

		// Setup routes
		for _, r := range cog.Routes {
			opts = append(opts, fx.Annotate(
				r,
				fx.As(new(api.Controller)),
				fx.ResultTags(`group:"routes"`),
			))
		}
	}
	return opts
}

func SetupAndRun() {
	fx.New(
		fx.Provide(
			setupLifeCycle,
			api.NewHTTPServer,
			fx.Annotate(
				api.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Provide(setupInternalDeps()...),
		fx.Provide(setupModuleConfigs()...),
		fx.Invoke(func(*AppInstance) {}),
	).Run()
}

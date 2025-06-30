package healthcheck

import "go-modular-sample/internal/shared/module"

func Config() module.ModuleConfig {
	return module.ModuleConfig{
		UseCases: UseCases(),
		Deps:     Deps(),
		Routes:   Routes(),
	}
}

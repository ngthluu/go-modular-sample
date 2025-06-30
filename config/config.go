package config

import (
	"go-modular-sample/internal/shared/db"
	"go-modular-sample/internal/shared/module"

	"go-modular-sample/internal/module/healthcheck"
)

func GetAllInternalDeps() map[any]any {
	return map[any]any{
		new(db.Db): db.NewPostgreSqlDb(),
	}
}

func GetAllModuleConfigs() []module.ModuleConfig {
	return []module.ModuleConfig{
		healthcheck.Config(),
	}
}

package module

type (
	Deps     = map[any]any
	Routes   = []any
	UseCases = []any
)

type ModuleConfig struct {
	Deps     Deps
	Routes   Routes
	UseCases UseCases
}

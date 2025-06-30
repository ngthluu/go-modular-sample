package api

import (
	"go-modular-sample/internal/shared/api"
	"net/http"
)

var _ api.Controller = (*HealthCheckController)(nil)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (ctl *HealthCheckController) Path() string {
	return "/healthcheck"
}

func (ctl *HealthCheckController) Setup(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func (ctl *HealthCheckController) Run(w http.ResponseWriter, r *http.Request) {
	api.ResponseJson(w, http.StatusOK, map[string]any{
		"success": true,
	})
}

package http

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/restful-api-demo/apps"
	"github.com/summingyu/restful-api-demo/apps/host"
)

var handler = &Handler{}

type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	if apps.HostService == nil {
		panic("dependence host service required")
	}
	h.svc = apps.HostService
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryGin(handler)
}

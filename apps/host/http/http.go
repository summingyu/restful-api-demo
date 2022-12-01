package http

import (
	"github.com/gin-gonic/gin"

	"github.com/summingyu/restful-api-demo/apps/host"
)

func NewHostHTTPHandler(svc host.Service) *Handler {
	return &Handler{
		svc: svc,
	}

}

type Handler struct {
	svc host.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

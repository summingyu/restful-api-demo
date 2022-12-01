package http

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"

	"github.com/summingyu/restful-api-demo/apps/host"
)

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	var err error
	ins, err = h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)

}

package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/summingyu/restful-api-demo/apps/host"
)

var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}

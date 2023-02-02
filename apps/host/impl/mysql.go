package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/summingyu/restful-api-demo/apps"
	"github.com/summingyu/restful-api-demo/apps/host"
	"github.com/summingyu/restful-api-demo/conf"
)

var impl = &HostServiceImpl{}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

func (i *HostServiceImpl) Config() {
	i.l = zap.L().Named(host.AppName)
	i.db = conf.C().MySQL.GetDB()

}
func (i *HostServiceImpl) Name() string {
	return host.AppName

}

func init() {
	apps.RegistryImpl(impl)
}

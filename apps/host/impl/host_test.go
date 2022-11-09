package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/summingyu/restful-api-demo/apps/host"
	"github.com/summingyu/restful-api-demo/apps/host/impl"
)

var (
	service host.Service
)

func TestCreateHost(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	_, _ = service.CreateHost(context.Background(), ins)

}

func init() {
	_ = zap.DevelopmentSetup()
	service = impl.NewHostServiceImpl()
}

package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"github.com/summingyu/restful-api-demo/apps/host"
	"github.com/summingyu/restful-api-demo/apps/host/impl"
	"github.com/summingyu/restful-api-demo/conf"
)

var (
	service host.Service
)

func TestCreateHost(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "杭州"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}

}

func init() {
	_ = zap.DevelopmentSetup()
	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}
	service = impl.NewHostServiceImpl()

}

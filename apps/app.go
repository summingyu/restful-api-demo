package apps

import (
	"fmt"

	"github.com/summingyu/restful-api-demo/apps/host"
)

// IOC 容器层: 管理所有的服务的实例

// 1. HostService 的实例必须注册过来, HostService才会有具体的实例, 服务启动时注册
// 2. HTTP 暴露模块, 依赖IOC里面的HostService
var (
	HostService host.Service
	svcs        = map[string]Service{}
)

func Registry(svc Service) {
	if _, ok := svcs[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))

	}
	svcs[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

func Init() {
	for _, v := range svcs {
		v.Config()
	}
}

type Service interface {
	Config()
	Name() string
}

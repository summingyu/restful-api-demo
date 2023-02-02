package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/summingyu/restful-api-demo/apps/host"
)

// IOC 容器层: 管理所有的服务的实例

// 1. HostService 的实例必须注册过来, HostService才会有具体的实例, 服务启动时注册
// 2. HTTP 暴露模块, 依赖IOC里面的HostService
var (
	HostService host.Service
	implApps    = map[string]ImplService{}
	ginApps     = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("impl service %s has registried", svc.Name()))

	}
	implApps[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

func RegistryGin(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("gin service %s has registried", svc.Name()))

	}
	ginApps[svc.Name()] = svc

}

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

func InitGin(r gin.IRouter) {
	for _, v := range ginApps {
		v.Config()
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

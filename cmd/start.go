package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/summingyu/restful-api-demo/apps"
	_ "github.com/summingyu/restful-api-demo/apps/all"
	"github.com/summingyu/restful-api-demo/apps/host/http"
	"github.com/summingyu/restful-api-demo/conf"
)

var (
	confType string
	confFile string
	confETCD string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = zap.DevelopmentSetup()
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}
		// service := impl.NewHostServiceImpl()
		apps.Init()
		api := http.NewHostHTTPHandler()
		api.Config()
		g := gin.Default()
		api.Registry(g)
		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)

}

package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/summingyu/restful-api-demo/apps"
	_ "github.com/summingyu/restful-api-demo/apps/all"
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
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}

		if err := loadGlobalLogger(); err != nil {
			return err
		}
		// service := impl.NewHostServiceImpl()
		apps.InitImpl()
		g := gin.Default()
		apps.InitGin(g)
		return g.Run(conf.C().App.HttpAddr())
	},
}

func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	// 根据config里面的日志配置, 来配置全局Logger对象
	lc := conf.C().Log
	// 设置日志级别
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()
	// 使用用户自定义配置, 替换默认配置
	zapConfig.Level = level
	zapConfig.Files.RotateOnStartup = false
	// 日志输出方式
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)

}

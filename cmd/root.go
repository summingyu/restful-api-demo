package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/summingyu/restful-api-demo/version"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo 后端API",
	Long:  "demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
}

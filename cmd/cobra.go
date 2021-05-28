package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.weimiaocaishang.com/weimiao/base_api/cmd/api_server"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/helper"
)

var rootCmd = &cobra.Command{
	Use:          "gitlab.weimiaocaishang.com/weimiao/base_api",
	Short:        "gitlab.weimiaocaishang.com/weimiao/base_api",
	SilenceUsage: true,
	Long:         `gitlab.weimiaocaishang.com/weimiao/base_api`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(helper.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + helper.Green(`gitlab.weimiaocaishang.com/weimiao/base_api `) + ` 可以使用 ` + helper.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api_server.ServerCmd)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yijieyu/go_basic_api/cmd/api_server"
	"github.com/yijieyu/go_basic_api/pkg/helper"
)

var rootCmd = &cobra.Command{
	Use:          "github.com/yijieyu/go_basic_api",
	Short:        "github.com/yijieyu/go_basic_api",
	SilenceUsage: true,
	Long:         `github.com/yijieyu/go_basic_api`,
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
	usageStr := `欢迎使用 ` + helper.Green(`github.com/yijieyu/go_basic_api `) + ` 可以使用 ` + helper.Red(`-h`) + ` 查看命令`
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

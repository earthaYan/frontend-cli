package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "custom-frontend-cli",
	Short:   "A frontend scaffold tool",
	Version: "0.0.1",
}
var createCmd = &cobra.Command{
	Use:   "create projectName",
	Short: "create a new project with certain project name",
	Args:  cobra.ExactArgs(1),
	Run:   createProject,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("命令执行失败")
		os.Exit(1)
	}
}

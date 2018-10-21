package cmd

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/controller"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var zhenlyChenCmd = &cobra.Command{
	Use:   "zhenlychen",
	Short: "test command",
	Run:   wrapper(controller.Other().Zhenlychen),
}

func init() {
	rootCmd.AddCommand(zhenlyChenCmd)
	zhenlyChenCmd.Flags().StringP("name", "n", "", "your name")
}

package cmd

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/controller"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a user",
	Run:   wrapper(controller.User().Register),
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("tel", "t", "", "user telephone")
}

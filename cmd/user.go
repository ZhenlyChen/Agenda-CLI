package cmd

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/controller"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var userCmd = &cobra.Command{
	Use: "user",
	Aliases: []string{"u"},
	Short: "User Command",
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Aliases: []string{"r"},
	Short: "register a user",
	Run:   wrapper(controller.User().Register),
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User log in",
	Run:   wrapper(controller.User().Login),
}

var logoutCmd  = &cobra.Command{
	Use:   "logout",
	Aliases: []string{"exit", "quit"},
	Short: "Log out",
	PreRun: func(cmd *cobra.Command, args []string) {
		controller.User().CheckLogin()
	},
	Run:   wrapper(controller.User().Logout),
}

var statusCmd = &cobra.Command{
	Use: "status",
	Short: "View the currently logged in user",
	Run: wrapper(controller.User().Status),
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all users",
	Run: wrapper(controller.User().List),
}

var DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "delete users",
	PreRun: func(cmd *cobra.Command, args []string) {
		controller.User().CheckLogin()
	},
	Run: wrapper(controller.User().Delete),
}

func init() {
	// 登陆命令
	loginCmd.Flags().StringP("user", "u", "", "username")
	loginCmd.Flags().StringP("password", "p", "", "user password")
	rootCmd.AddCommand(loginCmd)
	// 退出登陆
	rootCmd.AddCommand(logoutCmd)
	// 状态
	rootCmd.AddCommand(statusCmd)
	// 用户类命令
	rootCmd.AddCommand(userCmd)
	// 注册命令
	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("tel", "t", "", "user telephone")
	userCmd.AddCommand(registerCmd)
	// 查询命令
	userCmd.AddCommand(listCmd)
	// 删除命令
	userCmd.AddCommand(DeleteCmd)
}

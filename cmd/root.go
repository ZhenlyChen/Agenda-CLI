package cmd

import (
	"fmt"
	"os"

	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "A simple agenda management system",
	Long:  `Agenda can help you manage your agenda`,
}

// Execute 程序执行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./.agenda.yaml)")
}

// 初始化配置
func initConfig() {
	viper.SetDefault("DataBase.User", "./data/user.json")
	viper.SetDefault("DataBase.Meeting", "./data/meeting.json")
	viper.SetDefault("DataBase.Status", "./data/status.json")
	viper.SetDefault("LogFile", "./logs/agenda.txt")
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(".agenda")
	}
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "[Warning] Can't read config file in \""+viper.ConfigFileUsed()+"\", agenda will use default config.")
	}
	// 初始化数据库
	if err := model.InitDB(model.DataFile{
		User:    viper.GetString("DataBase.User"),
		Meeting: viper.GetString("DataBase.Meeting"),
		Status:  viper.GetString("DataBase.Status"),
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// 初始化日志系统
	if err := util.Log().Init(viper.GetString("LogFile")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

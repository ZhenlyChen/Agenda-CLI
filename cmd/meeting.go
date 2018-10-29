package cmd

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/controller"
	"github.com/spf13/cobra"
)

var meetingCmd = &cobra.Command{
	Use: "meeting",
	Aliases: []string{"m"},
	Short: "Meeting Command",
}

var meetingCreate = &cobra.Command{
	Use: "create",
	Short: "Create a meeting",
	Run: wrapper(controller.Meeting().Create),
}

var meetingAddParticipator = &cobra.Command{
	Use: "add",
	Short: "Add participator",
	Run: wrapper(controller.Meeting().AddParticipator),
}

var meetingRemoveParticipator = &cobra.Command{
	Use: "remove",
	Short: "Remove participator",
	Run: wrapper(controller.Meeting().RemoveParticipator),
}

var meetingQuery = &cobra.Command{
	Use: "query",
	Short: "Add participator",
	Aliases: []string{ "search"},
	Run: wrapper(controller.Meeting().Query),
}

var meetingDelete = &cobra.Command{
	Use: "delete",
	Short: "Delete Meeting by title",
	Run: wrapper(controller.Meeting().MeetingDelete),
}

var meetingQuit = &cobra.Command{
	Use: "quit",
	Short: "Quit Meeting by title",
	Run: wrapper(controller.Meeting().MeetingQuit),
}

var meetingClear = &cobra.Command{
	Use: "add",
	Short: "Add participator",
	Run: wrapper(controller.Meeting().Clear),
}


func init() {
	// 会议类命令
	rootCmd.AddCommand(meetingCmd)
	// 创建会议
	meetingCreate.Flags().StringP("title","t","","Meeting's title")
	meetingCreate.Flags().StringP("participator","p","","Meeting's participator")
	meetingCreate.Flags().StringP("start","s","","Meeting's start time")
	meetingCreate.Flags().StringP("end","e","","Meeting's end time")
	meetingCmd.AddCommand(meetingCreate)
	// 添加参与者
	meetingAddParticipator.Flags().StringP("title","t","","Meeting's title")
	meetingAddParticipator.Flags().StringP("participator","p","","Meeting's participator")
	meetingCmd.AddCommand(meetingAddParticipator)
	// 查询会议
	meetingQuery.Flags().StringP("start","s","","Query start time")
	meetingQuery.Flags().StringP("end","e","","Query end time")
	meetingCmd.AddCommand(meetingQuery)
}

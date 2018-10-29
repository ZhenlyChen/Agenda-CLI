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
	Use: "add",
	Short: "Add participator",
	Run: wrapper(controller.Meeting().RemoveParticipator),
}

var meetingQuery = &cobra.Command{
	Use: "add",
	Short: "Add participator",
	Run: wrapper(controller.Meeting().Query),
}

var meetingDelete = &cobra.Command{
	Use: "add",
	Short: "Add participator",
	Run: wrapper(controller.Meeting().MeetingDelete),
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
}

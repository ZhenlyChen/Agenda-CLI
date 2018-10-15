package controller

import (
	"github.com/spf13/cobra"
)

var ctrl *ctrlManger

func init() {
	ctrl = new(ctrlManger)
}

type ctrlManger struct {
	cmd  *cobra.Command
	args []string
}

// User 用户控制层
func User() UserInterface { return ctrl }

// BindData 绑定数据
func BindData(c *cobra.Command, a []string) { ctrl.bindData(c, a) }
func (c *ctrlManger) bindData(cmd *cobra.Command, args []string) {
	ctrl.cmd = cmd
	ctrl.args = args
}

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
func User() UserInterface { return ctrl	}

// BindData 绑定数据
func BindData(c *cobra.Command, a []string) { ctrl.bindData(c, a) }
func (ctrl *ctrlManger) bindData(c *cobra.Command, a []string) {
	ctrl.cmd = c
	ctrl.args = a
}

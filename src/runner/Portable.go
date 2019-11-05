package runner

import (
	"CSAE-Launcher-Plugin/src/common/utils"
	"fmt"
	"github.com/lxn/walk"
)

func Portable() {
	fmt.Println("注意: 此窗口非常重要，请不要关闭!")
	fmt.Println("提示: 本模式推荐网吧用户使用，非网吧用户请使用安装程序进行安装。")

	// 判断配置文件是否存在
	cfg := utils.ReadConfig()

	if !utils.PathExists(cfg.CSAE.Full) {
		walk.MsgBox(nil, "提示", "配置文件不存在或不正确，请重新选择游戏路径\n请在选择文件后重新启动本程序", walk.MsgBoxIconWarning)
		ChooseFile()
	}
	go Service()
}

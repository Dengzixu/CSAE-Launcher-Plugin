package runner

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/errorEx"
	"CSAE-Launcher-Plugin/src/common/utils"
	"github.com/lxn/walk"
	"go.uber.org/zap"
	"os"
)

func ChooseFile() {
	path, err := utils.ChooseFile()

	if nil != err {
		switch err.(*errorEx.Error).Code {
		case errorEx.ChooseFailed:
			walk.MsgBox(nil, "CSAE 启动器插件", "选择文件初始化失败，请联系管理员", walk.MsgBoxIconError)
		case errorEx.ConfigFileWriteFail:
			walk.MsgBox(nil, "CSAE 启动器插件", "无法创建配置文件，请联系管理员", walk.MsgBoxIconError)
		case errorEx.ChooseCancel:
			walk.MsgBox(nil, "CSAE 启动器插件", "尚未选择游戏文件，可能无法正常启动游戏", walk.MsgBoxIconWarning)
		}
		Logs.G.Errorw("文件选择失败", zap.Error(err))
	} else {
		Logs.G.Infow("文件选择完毕", zap.String("path", path))
	}

	os.Exit(0)
}

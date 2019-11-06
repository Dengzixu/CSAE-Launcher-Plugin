package Init

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"go.uber.org/zap"
)

func CheckUpdate() {
	Logs.G.Info("Check Update...")
	Logs.G.Warnw("Check Update Finish", zap.String("result", "not implemented"))
	//log.WithField("component", "Check Update").Info("检查更新...")
	//log.WithField("component", "Check Update").Warn("检查更新尚未完成")
}

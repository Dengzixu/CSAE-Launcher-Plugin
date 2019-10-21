package Init

import (
	log "github.com/sirupsen/logrus"
)

func CheckUpdate() {
	log.WithField("component", "Check Update").Info("检查更新...")
	log.WithField("component", "Check Update").Warn("检查更新尚未完成")
}

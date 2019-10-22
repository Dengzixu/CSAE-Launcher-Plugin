package Init

import (
	"CSAELauncherPlugin/common/utils"
	log "github.com/sirupsen/logrus"
)

func Config() {
	log.WithField("component", "Config").Info("读取配置文件...")
	utils.ReadConfig()
}

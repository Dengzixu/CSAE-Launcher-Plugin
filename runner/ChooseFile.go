package runner

import (
	"CSAELauncherPlugin/common/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func ChooseFile() {
	path, err := utils.ChooseFile()

	if nil != err {
		log.WithField("component", "Choose File").Error(err)
	}

	log.WithField("component", "Choose File").Info(path)

	os.Exit(0)
}

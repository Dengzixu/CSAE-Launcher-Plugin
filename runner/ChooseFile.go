package runner

import (
	"CSAELauncherPlugin/common/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func ChooseFile() {
	path, _ := utils.ChooseFile()

	log.WithField("component", "Choose File").Info("CSAE: ", path)

	os.Exit(0)
}

package runner

import (
	"CSAELauncherPlugin/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func ChooseFile() {
	path, _ := utils.ChooseFile()

	log.WithField("component", "Choose File").Info("CSAE: ", path)

	os.Exit(0)
}

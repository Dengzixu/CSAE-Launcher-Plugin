package runner

import (
	"CSAELauncherPlugin/utils"
	"fmt"
	"os"
)

func ChooseFile() {
	path, _ := utils.ChooseFile()

	fmt.Println(path)

	os.Exit(0)
}

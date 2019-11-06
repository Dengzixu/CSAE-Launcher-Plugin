package Init

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/utils"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func Logger() {
	Logs.Init()

	//log.SetOutput(os.Stdout)
	log.SetOutput(colorable.NewColorableStdout())
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&utils.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category", "req"},
	})

	//file, err := os.OpenFile("e://log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////defer file.Close()
	//log.SetOutput(file)
}

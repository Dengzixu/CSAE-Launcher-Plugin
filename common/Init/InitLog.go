package Init

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Logger() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category", "req"},
	})

	file, err := os.OpenFile("e://log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()
	log.SetOutput(file)
}

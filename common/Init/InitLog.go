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
}

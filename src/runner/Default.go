package runner

import (
	"flag"
	"os"
)

func Default() {
	flag.PrintDefaults()
	os.Exit(0)
}

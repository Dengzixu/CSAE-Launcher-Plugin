package Init

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func First() {
	localDir, _ := filepath.Split(os.Args[0])

	if l, _ := pathExists(localDir + "\\init.lock"); !l {
		log.Info("第一次运行, 执行初始化操作。")
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

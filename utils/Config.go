package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

const iniFile = "\\config.ini"

type CSAE struct {
	Exe  string `ini:"exe" comment:"CSAE 可执行文件名"`
	Dir  string `ini:"dir" comment:"CSAE 可执行文件所在路径"`
	Full string `ini:"full" comment:"完整路径"`
}

type Security struct {
	Version int `ini:"version"`
}

type ConfigV1 struct {
	Version  int `ini:"version" comment:"此行为配置文件版本，与程序版本无关，请勿修改"`
	CSAE     CSAE
	Security Security
}

func ReadConfig() *ConfigV1 {
	config := &ConfigV1{}

	cfg, _ := ini.LoadSources(ini.LoadOptions{
		IgnoreContinuation: true,
	}, configPath())

	_ = cfg.MapTo(config)

	return config
}

func WritePath(fullPath string) error {
	currentDir := configPath()

	cfg, err := ini.Load(configPath())

	csaeDir, csaeExe := filepath.Split(fullPath)

	fmt.Println("configPath: ", configPath())
	if nil != err {
		return err
	}

	cfg.Section("CSAE").Key("exe").SetValue(csaeExe)
	cfg.Section("CSAE").Key("dir").SetValue(csaeDir)
	cfg.Section("CSAE").Key("full").SetValue(fullPath)

	if cfg.SaveTo(currentDir) != nil {
		return err
	}
	return nil
}

func configPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return dir + iniFile
}

func CreateDefaultConfig() {
	config := &ConfigV1{
		Version: 1,
		CSAE: CSAE{
			Exe:  "",
			Dir:  "",
			Full: "",
		},
		Security: Security{Version: 1},
	}

	cfg := ini.Empty()

	_ = ini.ReflectFrom(cfg, config)

	_ = cfg.SaveTo(configPath())
}

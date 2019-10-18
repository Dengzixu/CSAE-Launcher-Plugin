package utils

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

const iniFile = "\\config.ini"

type ConfigV1 struct {
	Version int `ini:"version"`

	CSAE struct {
		Exe  string `ini:"exe"`
		Dir  string `ini:"dir"`
		Full string `ini:"full"`
	}
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
	dir, exe := filepath.Split(fullPath)

	cfg, err := ini.Load(configPath())
	if nil != err {
		return err
	}

	cfg.Section("CSAE").Key("exe").SetValue(exe)
	cfg.Section("CSAE").Key("dir").SetValue(dir)
	cfg.Section("CSAE").Key("full").SetValue(fullPath)

	if cfg.SaveTo(configPath()) != nil {
		return err
	}
	return nil
}

func configPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir + iniFile
}

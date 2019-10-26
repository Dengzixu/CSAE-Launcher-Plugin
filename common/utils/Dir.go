package utils

import (
	"os"
	"path/filepath"
)

const (
	dBase   = "\\.csae"
	dConfig = "\\config"
	dSSL    = "\\ssl"
)

const (
	tBase   = "tBase"
	tConfig = "tConfig"
	tSSL    = "tSSL"
)

func getDir(t string) (destDir string) {
	userHomeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	baseDir := userHomeDir + dBase

	switch t {
	case tBase:
		return baseDir
	case tConfig:
		return baseDir + dConfig
	case tSSL:
		return baseDir + dSSL
	}

	return ""
}

func GetBaseDir() string {
	return getDir(tBase)
}

func GetConfigDir() string {
	return getDir(tConfig)
}

func GetSSLDir() string {
	return getDir(tSSL)
}

func CreateDateDir() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// 创建基本目录
	_ = os.MkdirAll(dir+dBase, 0644)
	// 创建配置目录
	_ = os.MkdirAll(dir+dBase+dConfig, 0644)
	// 创建安全目录
	_ = os.MkdirAll(dir+dBase+dSSL, 0644)
}

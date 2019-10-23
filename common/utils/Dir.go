package utils

import (
	"os"
)

const (
	dBase   = "\\.csae\\launcher"
	dConfig = "\\config"
	dSSL    = "\\ssl"
)

const (
	tBase   = "tBase"
	tConfig = "tConfig"
	tSSL    = "tSSL"
)

func getDir(t string) (destDir string) {
	userHomeDir, _ := os.UserHomeDir()
	
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
	dir, _ := os.UserHomeDir()
	// 创建基本目录
	_ = os.MkdirAll(dir+dBase, 0644)
	// 创建配置目录
	_ = os.MkdirAll(dir+dBase+dConfig, 0644)
	// 创建安全目录
	_ = os.MkdirAll(dir+dBase+dSSL, 0644)
}

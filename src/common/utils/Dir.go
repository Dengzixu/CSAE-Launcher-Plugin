package utils

import (
	"os"
	"path/filepath"
)

const (
	dBase   = "\\.csae"
	dConfig = "\\config"
	dSSL    = "\\ssl"
	dLog    = "\\log"
)

const (
	tBase   = "tBase"
	tConfig = "tConfig"
	tSSL    = "tSSL"
	tLog    = "tLog"
)

const defaultFilePermission = 0644

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
	case tLog:
		return baseDir + dLog
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

func GetLogDir() string {
	return getDir(tLog)
}

func CreateDateDir() {
	// 创建基本目录
	_ = os.MkdirAll(GetBaseDir(), defaultFilePermission)
	// 创建配置目录
	_ = os.MkdirAll(GetConfigDir(), defaultFilePermission)
	// 创建安全目录
	_ = os.MkdirAll(GetSSLDir(), defaultFilePermission)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else if os.IsNotExist(err) {
		return false
	}

	return true
}

package Init

import (
	"CSAE-Launcher-Plugin/common/utils"
	"CSAE-Launcher-Plugin/entity"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"os"
)

const lockFileName = "\\.lock"

func First() {
	if utils.PathExists(utils.GetBaseDir() + lockFileName) {
		return
	}

	// 创建配置文件夹
	utils.CreateDateDir()

	log.WithField("component", "Init").Info("首次运行, 执行初始化操作...")

	// 创建配置文件
	log.WithField("component", "Init").Info("创建配置文件...")
	utils.CreateDefaultConfig()

	// 创建证书
	log.WithField("component", "Init").Info("创建安全证书...")
	createSSL()

	writeLockFile()
}

func createSSL() {
	securityConfig, err := utils.GetSecurityConfig()

	// 如果无法获取 则采用内置的证书
	if nil != err {
		securityConfig = entity.InnerSecurityConfig()
	}

	// 创建文件夹
	_ = os.Mkdir(utils.GetSSLDir(), 0644)

	// 创建证书
	certFile, _ := os.OpenFile(utils.GetSSLDir()+"\\certificate.crt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer certFile.Close()

	decodeCert, _ := base64.StdEncoding.DecodeString(securityConfig.TlsCertificate.Certificate)

	_, _ = certFile.Write(decodeCert)

	// 创建密钥
	keyFile, _ := os.OpenFile(utils.GetSSLDir()+"\\key.pem", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer keyFile.Close()

	keyCert, _ := base64.StdEncoding.DecodeString(securityConfig.TlsCertificate.Key)

	_, _ = keyFile.Write(keyCert)
}

func writeLockFile() {
	lockFile, _ := os.OpenFile(utils.GetBaseDir()+lockFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer lockFile.Close()

	_, _ = lockFile.WriteString("")
}

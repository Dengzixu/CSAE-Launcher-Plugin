package Init

import (
	"CSAELauncherPlugin/common/utils"
	"CSAELauncherPlugin/entity"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const lockFileName = "\\.lock"

func First() {
	localDir, _ := filepath.Split(os.Args[0])

	if pathExists(localDir + lockFileName) {
		return
	}
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
	localDir, _ := filepath.Split(os.Args[0])

	securityConfig, err := utils.GetSecurityConfig()

	// 如果无法获取 则采用内置的证书
	if nil != err {
		securityConfig = entity.InnerSecurityConfig()
	}

	// 创建文件夹
	_ = os.Mkdir(localDir+"\\ssl", 0644)

	// 创建证书
	certFile, _ := os.OpenFile(localDir+"\\ssl\\certificate.crt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer certFile.Close()

	decodeCert, _ := base64.StdEncoding.DecodeString(securityConfig.TlsCertificate.Certificate)

	_, _ = certFile.Write(decodeCert)

	// 创建密钥
	keyFile, _ := os.OpenFile(localDir+"\\ssl\\key.pem", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer keyFile.Close()

	keyCert, _ := base64.StdEncoding.DecodeString(securityConfig.TlsCertificate.Key)

	_, _ = keyFile.Write(keyCert)
}

func writeLockFile() {
	localDir, _ := filepath.Split(os.Args[0])

	lockFile, _ := os.OpenFile(localDir+lockFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer lockFile.Close()

	_, _ = lockFile.WriteString("")
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else if os.IsNotExist(err) {
		return false
	}

	return true
}

package utils

import (
	"CSAELauncherPlugin/common"
	"CSAELauncherPlugin/entity"
	"fmt"
	"os"
	"os/exec"
)

func LaunchGameOffline(path string, param string) (int, error) {
	cmd := exec.Command(path, param)

	if err := cmd.Start(); err != nil {
		return msg.ErrLaunchFail, err
	}
	return msg.Success, nil
}

func LaunchGameOnline(path string, param string, host string, password string) (int, error) {
	if "" != password {
		param += " +password \"" + password + "\""
	}
	param += " +connect " + host

	cmd := exec.Command(path, param)

	if err := cmd.Start(); err != nil {
		return msg.ErrLaunchFail, err
	}
	return msg.Success, nil
}

func LaunchGame(config *entity.LaunchConfig) (int, error) {
	sysConfig := ReadConfig()

	// HOST 为空的话 就只启动游戏
	if "" == config.Host {
		return LaunchGameOffline(sysConfig.CSAE.Full, config.Option)
	} else {
		pwdResp, err := GetPassword(config.Host, config.Token)
		if nil == err {
			switch pwdResp.Code {
			case 403:
				return msg.ErrApiPermission, fmt.Errorf("未登录或登录过期")
			default:
				return msg.ErrApi, fmt.Errorf("未知错误")
			}
		}

		userInfoResp, err2 := GetUserInfo(config.Token)
		if nil == err2 {
			switch userInfoResp.Code {
			case 403:
				return msg.ErrApiPermission, fmt.Errorf("未登录或登录过期")
			default:
				return msg.ErrApi, fmt.Errorf("未知错误")
			}
		}

		if _, err = setConfig(sysConfig.CSAE.Dir, userInfoResp.Data.Nickname); nil != err {
			return msg.ErrWriteConfig, err
		}

		return LaunchGameOnline(sysConfig.CSAE.Full, config.Option, config.Host, pwdResp.Data.ServerPass)
	}
}

func setConfig(dir string, username string) (int, error) {
	userConfig := `cl_cmdbackup "2"
cl_cmdrate "100"
cl_timeout "300"
cl_updaterate "100"
console "1.0"
fps_max "100"
rate "25000"
`

	userConfig += "name " + username + "\n" +
		"alias name"

	userConfigFile, err := os.OpenFile(dir+"\\cstrike_schinese\\userconfig.cfg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if nil != err {
		return msg.ErrWriteConfig, err
	}
	defer userConfigFile.Close()

	userConfigByte := []byte(userConfig)

	if _, err := userConfigFile.Write(userConfigByte); err != nil {
		return msg.ErrWriteConfig, err
	}

	return msg.Success, nil
}

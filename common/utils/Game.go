package utils

import (
	"CSAELauncherPlugin/common/global"
	"CSAELauncherPlugin/common/msg"
	"CSAELauncherPlugin/entity"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func LaunchGameOffline(path string, param string) (int, error) {

	if global.IsService {
		dir, _ := filepath.Split(path)
		if err := StartProcessAsCurrentUser(path, param, dir, false); nil != err {
			return msg.ErrLaunchFail, err
		}
	} else {
		cmd := exec.Command(path, param)
		if err := cmd.Start(); err != nil {
			return msg.ErrLaunchFail, err
		}
	}
	return msg.Success, nil

	//cmd := exec.Command(path, param)
	//
	//if err := cmd.Start(); err != nil {
	//	return msg.ErrLaunchFail, err
	//}
	//return msg.Success, nil
}

func LaunchGameOnline(path string, param string, host string, password string) (int, error) {
	if "" != password {
		param += " +password \"" + password + "\""
	}
	param += " +connect " + host

	if global.IsService {
		dir, _ := filepath.Split(path)

		_ = StartProcessAsCurrentUser(path, param, dir, false)

		if err := StartProcessAsCurrentUser(path, param, dir, false); nil != err {
			return msg.ErrLaunchFail, err
		}
	} else {
		cmd := exec.Command(path, param)
		if err := cmd.Start(); err != nil {
			return msg.ErrLaunchFail, err
		}
	}
	return msg.Success, nil
}

func LaunchGame(config *entity.LaunchConfig) (int, error) {
	sysConfig := ReadConfig()

	// HOST 为空的话 就只启动游戏
	if "" == config.Host {
		return LaunchGameOffline(sysConfig.CSAE.Full, config.Option)
	} else {
		// 密码方式
		var serverPassword string
		switch config.Password {
		// 启动器启动
		case "#LAUNCHER#":
			pwdResp, err := GetPassword(config.Host, config.Token)
			if nil != err {
				return msg.ErrApiNetwork, fmt.Errorf(msg.GetMsg(msg.ErrApiNetwork))
			} else if 0 != pwdResp.Code {
				return msg.ErrApiPermission, fmt.Errorf(msg.GetMsg(msg.ErrApiPermission))
			} else {
				serverPassword = pwdResp.Data.ServerPass
			}
		case "#NOPASSWORD#":
			serverPassword = ""
		default:
			serverPassword = config.Password
		}

		// 获取用户昵称
		userInfoResp, err2 := GetUserInfo(config.Token)
		if nil != err2 {
			return msg.ErrApiNetwork, fmt.Errorf(msg.GetMsg(msg.ErrApiNetwork))
		} else if 0 != userInfoResp.Code {
			return msg.ErrApiPermission, fmt.Errorf(msg.GetMsg(msg.ErrApiPermission))
		}

		if _, err3 := setConfig(sysConfig.CSAE.Dir, userInfoResp.Data.Nickname); nil != err3 {
			return msg.ErrWriteConfig, err3
		}

		return LaunchGameOnline(sysConfig.CSAE.Full, config.Option, config.Host, serverPassword)
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

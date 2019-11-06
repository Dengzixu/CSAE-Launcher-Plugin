package utils

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/errorEx"
	"CSAE-Launcher-Plugin/src/common/global"
	"CSAE-Launcher-Plugin/src/entity"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
)

func LaunchGame(config *entity.LaunchConfig) error {
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
				Logs.G.Errorw("Get password failed", zap.Error(err))
				return errorEx.New(errorEx.ApiNetWorkError)
			} else if 0 != pwdResp.Code {
				Logs.G.Errorw("Get password failed", "API response error", zap.Int("code", pwdResp.Code))
				return errorEx.New(errorEx.ApiPermissionError)
			} else {
				serverPassword = pwdResp.Data.ServerPass
			}
		case "#NOPASSWORD#":
			serverPassword = ""
		default:
			serverPassword = config.Password
		}

		// 获取用户昵称
		userInfoResp, err := GetUserInfo(config.Token)
		if nil != err {
			Logs.G.Errorw("Get username failed", zap.Error(err))
			return errorEx.New(errorEx.ApiNetWorkError)
		} else if 0 != userInfoResp.Code {
			Logs.G.Errorw("Get username failed", "API response error", zap.Int("code", userInfoResp.Code))
			return errorEx.New(errorEx.ApiPermissionError)
		}

		if err := setConfig(sysConfig.CSAE.Dir, userInfoResp.Data.Nickname); nil != err {
			return err
		}

		return LaunchGameOnline(sysConfig.CSAE.Full, config.Option, config.Host, serverPassword)
	}
}

func LaunchGameOffline(path string, param string) error {
	err := run(path, param)
	return err
}

func LaunchGameOnline(path string, param string, host string, password string) error {
	if "" != password {
		param += " +password " + password
	}
	param += " +connect " + host

	err := run(path, param)
	return err
}

func run(path string, param string) error {
	Logs.G.Debug("Launch game", zap.String("path", path))
	if global.IsService {
		dir, _ := filepath.Split(path)
		if err := StartProcessAsCurrentUser(path, param, dir, false); nil != err {
			Logs.G.Errorw("Launch game failed", zap.Bool("isService", global.IsService))
			return errorEx.New(errorEx.LaunchGameFailed)
		}
	} else {
		cmd := exec.Command(path, param)
		if err := cmd.Start(); err != nil {
			Logs.G.Errorw("Launch game failed", zap.Bool("isService", global.IsService))
			return errorEx.New(errorEx.LaunchGameFailed)
		}
	}
	return nil
}

func setConfig(dir string, username string) error {
	userConfig := `cl_cmdbackup "2"
cl_cmdrate "100"
cl_timeout "300"
cl_updaterate "100"
console "1.0"
fps_max "100"
rate "25000"
name "%s"
alias name
`

	userConfig = fmt.Sprintf(userConfig, username)

	userConfigFile, err := os.OpenFile(dir+"\\cstrike_schinese\\userconfig.cfg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if nil != err {
		return errorEx.New(errorEx.GameConfigFileLoadFail)
	}
	defer userConfigFile.Close()

	userConfigByte := []byte(userConfig)

	if _, err := userConfigFile.Write(userConfigByte); err != nil {
		return errorEx.New(errorEx.GameConfigFileWriteFail)
	}

	return nil
}

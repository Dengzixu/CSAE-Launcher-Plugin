package utils

import (
	"CSAELauncherPlugin/common"
	"CSAELauncherPlugin/entity"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
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
	// HOST 为空的话 就只启动游戏
	if "" == config.Host {
		return LaunchGameOffline(config.Path, config.Option)
	} else {

		pwdResp, err := GetPassword(config.Host, config.Token)
		if nil != err {
			return msg.ErrApi, errors.New("API 返回错误")
		}

		userInfoResp, err := GetUserInfo(config.Token)
		if nil != err {
			return msg.ErrApi, errors.New("API 返回错误")
		}

		if _, err = setConfig(config.Path, userInfoResp.Data.Nickname); nil != err {
			return msg.ErrWriteConfig, err
		}

		return LaunchGameOnline(config.Path, config.Option, config.Host, pwdResp.Data.ServerPass)
	}
}

func setConfig(path string, username string) (int, error) {
	path, _ = filepath.Split(path)

	data := "cl_cmdbackup \"2\"\ncl_cmdrate \"100\"\ncl_timeout \"300\"\ncl_updaterate \"100\"\nconsole \"1.0\"\nfps_max \"100\"\nrate \"25000\"\n"

	data += "name " + username + "\n" +
		"alias name"

	fileObj, err := os.OpenFile(path+"\\cstrike_schinese\\userconfig.cfg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if nil != err {
		return msg.ErrWriteConfig, err
	}
	defer fileObj.Close()
	contents := []byte(data)
	if _, err := fileObj.Write(contents); err != nil {
		return msg.ErrWriteConfig, err
	}

	return msg.Success, nil
}

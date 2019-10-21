package utils

import (
	"fmt"
	"github.com/lxn/walk"
)

func ChooseFile() (string, error) {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择CSAE程序位置"
	dlg.Filter = "CSAE程序 (csae.exe,hl.exe,csae_master.exe,hl_master.exe)|csae.exe;hl.exe;csae_master.exe;hl_master.exe|"

	if ok, err := dlg.ShowOpen(nil); err != nil {
		return "", err
	} else if !ok {
		return "选择文件被取消", nil
	}

	if err := WritePath(dlg.FilePath); nil != err {
		fmt.Println(err)

		return "写入配置文件失败", nil
	}

	return dlg.FilePath, nil
}

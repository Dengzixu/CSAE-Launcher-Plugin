package utils

import (
	"errors"
	"github.com/lxn/walk"
)

func ChooseFile() (string, error) {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择CSAE程序位置"
	dlg.Filter = "CSAE程序 (csae.exe,hl.exe,csae_master.exe,hl_master.exe)|csae.exe;hl.exe;csae_master.exe;hl_master.exe|"
	dlg.ShowReadOnlyCB = true

	if ok, err := dlg.ShowOpen(nil); err != nil {
		return "", err
	} else if !ok {
		return "选择文件被取消", errors.New("choose file was be cancel")
	}

	if err := WriteCSAEPath(dlg.FilePath); nil != err {
		return "写入配置文件失败", err
	}

	return dlg.FilePath, nil
}

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
		walk.MsgBox(nil, "错误", "无法创建配置文件", walk.MsgBoxIconError)
	} else if !ok {
		walk.MsgBox(nil, "警告", "无法创建配置文件", walk.MsgBoxIconError)
		return "选择文件被取消", errors.New("choose file was be cancel")
	}

	if err := WriteCSAEPath(dlg.FilePath); nil != err {
		walk.MsgBox(nil, "错误", "无法创建配置文件", walk.MsgBoxIconError)
		return "写入配置文件失败", err
	}

	return dlg.FilePath, nil
}

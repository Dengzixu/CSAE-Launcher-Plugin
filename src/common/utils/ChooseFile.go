package utils

import (
	"CSAE-Launcher-Plugin/src/common/errorEx"
	"github.com/lxn/walk"
)

func ChooseFile() (string, error) {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择CSAE程序位置"
	dlg.Filter = "CSAE程序 (csae.exe,hl.exe,csae_master.exe,hl_master.exe)|csae.exe;hl.exe;csae_master.exe;hl_master.exe|"
	dlg.ShowReadOnlyCB = true

	if ok, err := dlg.ShowOpen(nil); err != nil {
		return "", errorEx.New(errorEx.ChooseFailed)
	} else if !ok {
		return "", errorEx.New(errorEx.ChooseCancel)
	}

	if err := WriteCSAEPath(dlg.FilePath); nil != err {
		return "", errorEx.New(errorEx.ConfigFileWriteFail)
	}

	return dlg.FilePath, nil
}

package utils

import (
	"github.com/lxn/walk"
)

func ChooseFile() (string, error) {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择CSAE程序位置"
	dlg.Filter = "CSAE程序 (csae.exe,hl.exe,csae_master.exe,hl_master.exe)|csae.exe;hl.exe;csae_master.exe;hl_master.exe|"

	if ok, err := dlg.ShowOpen(nil); err != nil {
		return "", err
	} else if !ok {
		return "cancel", nil
	}
	return dlg.FilePath, nil
}

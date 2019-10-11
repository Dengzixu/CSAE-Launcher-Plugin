package utils

import (
    "os/exec"
)

func RunGameOffline(path string, param string) bool {
    cmd := exec.Command(path, param)

    err := cmd.Start()

    if nil != err {
        return false
    }

    return true
}

func RunGameOnline(path string, param string, host string, password string) bool {
    return false
}

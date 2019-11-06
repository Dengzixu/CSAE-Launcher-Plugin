package entity

import (
	"CSAE-Launcher-Plugin/src/common/errorEx"
	"CSAE-Launcher-Plugin/src/common/msg"
)

func RespBody(code uint32, success bool, data interface{}) map[string]interface{} {

	r := make(map[string]interface{})
	r["code"] = code
	r["success"] = success
	//r["message"] = msg.GetMsg(code)
	r["message"] = errorEx.GetMsg(code)
	r["data"] = data

	return r
}

func RespBodySuccess() map[string]interface{} {
	return RespBody(msg.Success, true, nil)
}

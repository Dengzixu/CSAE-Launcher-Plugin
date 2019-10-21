package entity

import (
	msg2 "CSAELauncherPlugin/common/msg"
)

func RespBody(code int, success bool, data interface{}) map[string]interface{} {

	r := make(map[string]interface{})
	r["code"] = code
	r["success"] = success
	r["message"] = msg2.GetMsg(code)
	r["data"] = data

	return r
}

func RespBodySuccess() map[string]interface{} {
	return RespBody(msg2.Success, true, nil)
}

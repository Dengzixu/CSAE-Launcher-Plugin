package entity

import msg "CSAELauncherPlugin/common"

func RespBody(code int, success bool, data interface{}) map[string]interface{} {

	r := make(map[string]interface{})
	r["code"] = code
	r["success"] = success
	r["message"] = msg.GetMsg(code)
	r["data"] = data

	return r
}

func RespBodySuccess() map[string]interface{} {
	return RespBody(msg.Success, true, nil)
}

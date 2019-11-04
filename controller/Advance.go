package controller

import (
	msg2 "CSAE-Launcher-Plugin/common/msg"
	"CSAE-Launcher-Plugin/common/utils"
	"CSAE-Launcher-Plugin/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LaunchController(context *gin.Context) {
	launchParam := &entity.LaunchConfig{}

	if err := context.BindJSON(&launchParam); nil != err {
		context.JSON(http.StatusBadRequest, entity.RespBody(msg2.ErrApiParam, false, nil))
	}

	if rCode, err := utils.LaunchGame(launchParam); nil != err {
		context.JSON(http.StatusBadGateway, entity.RespBody(rCode, false, nil))
		return
	}

	context.JSON(http.StatusOK, entity.RespBodySuccess())
}

func ChooseFileController(context *gin.Context) {
	context.String(http.StatusOK, "功能弃用")
	return
}

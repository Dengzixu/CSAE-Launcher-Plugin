package controller

import (
	msg "CSAELauncherPlugin/common"
	"CSAELauncherPlugin/entity"
	"CSAELauncherPlugin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LaunchController(context *gin.Context) {
	launchParam := &entity.LaunchConfig{}

	if err := context.BindJSON(&launchParam); nil != err {
		context.JSON(http.StatusBadRequest, entity.RespBody(msg.ErrApiParam, false, nil))
	}

	// 判断路径是否为空
	if "" == launchParam.Path {
		context.JSON(http.StatusBadRequest, entity.RespBody(msg.ErrLaunchPath, false, nil))
		return
	}

	resCode, err := utils.LaunchGame(launchParam)

	if nil != err {
		context.JSON(http.StatusInternalServerError, entity.RespBody(resCode, false, nil))
		return
	}

	context.JSON(http.StatusOK, entity.RespBodySuccess())
}

func ChooseFileController(context *gin.Context) {
	if filePath, err := utils.ChooseFile(); nil != err {
		context.JSON(http.StatusOK, entity.RespBody(msg.ErrChooseFail, false, nil))
	} else if filePath == "cancel" {
		context.JSON(http.StatusOK, entity.RespBody(msg.ErrChooseCancel, false, nil))
	} else {
		context.JSON(http.StatusOK, entity.RespBody(msg.Success, true, filePath))
	}
}

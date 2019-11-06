package controller

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/errorEx"
	"CSAE-Launcher-Plugin/src/common/msg"
	"CSAE-Launcher-Plugin/src/common/utils"
	"CSAE-Launcher-Plugin/src/entity"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func LaunchController(context *gin.Context) {
	launchParam := &entity.LaunchConfig{}

	if err := context.BindJSON(&launchParam); nil != err {
		Logs.G.Errorw("Can not resolve request", zap.Error(err))
		context.JSON(http.StatusBadRequest, entity.RespBody(msg.ErrApiParam, false, nil))
		return
	}

	if err := utils.LaunchGame(launchParam); nil != err {
		Logs.G.Errorw("Launch game failed", zap.Error(err))
		code := err.(*errorEx.Error).Code
		var httpStatusCode int = http.StatusInternalServerError

		switch code {
		case errorEx.ConfigFileLoadFail, errorEx.GameConfigFileWriteFail:
			httpStatusCode = http.StatusBadGateway
		case errorEx.ApiNetWorkError, errorEx.ApiPermissionError:
			httpStatusCode = http.StatusBadGateway
		}

		context.JSON(httpStatusCode, entity.RespBody(code, false, nil))
		return
	}

	context.JSON(http.StatusOK, entity.RespBodySuccess())
}

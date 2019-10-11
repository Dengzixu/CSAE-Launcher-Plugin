package controller

import (
    "CSAELauncherPlugin/entity"
    "CSAELauncherPlugin/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func LaunchController(context *gin.Context) {
    launchParam := &entity.LaunchParam{}

    err := context.BindJSON(&launchParam)

    if nil != err {
        context.String(http.StatusBadRequest, "请求参数错误")
    }

    utils.GetPassword(launchParam.Host, launchParam.Token)

    // 判断路径是否为空
    if "" == launchParam.Path {
        context.String(http.StatusBadRequest, "路径不能为空")
        return
    }

    // 如果没有HOST就启动游戏
    if "" == launchParam.Host {
        if !utils.RunGameOffline(launchParam.Path, launchParam.Param) {
            context.String(http.StatusInternalServerError, "启动游戏失败")
        }
        context.String(http.StatusOK, "成功")
        return
    } else {

        if !utils.RunGameOnline(launchParam.Path, launchParam.Param, launchParam.Host, "password") {
            context.String(http.StatusInternalServerError, "启动游戏失败")
        }
        context.String(http.StatusOK, "成功")
        return
    }
}

func ChooseFileController(context *gin.Context) {
    context.String(http.StatusOK, utils.ChooseFile())
}

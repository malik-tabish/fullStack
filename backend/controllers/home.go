package controllers

import (
	"gin-example/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, utils.NewMessage(http.StatusOK, nil, "Status Ok", "Welcome to my first application"))
}

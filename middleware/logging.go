package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-gin-todolist/pkg/logging"
)

func LoggerToFile(ctx *gin.Context) {
	logging.Logger.WithFields(logrus.Fields{
		"status_code": ctx.Writer.Status(),
		"client_ip": ctx.ClientIP(),
		"req_method": ctx.Request.Method,
		"req_uri": ctx.Request.RequestURI,
	}).Info()

	ctx.Next()
}
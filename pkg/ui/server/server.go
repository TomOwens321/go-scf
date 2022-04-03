// based loosly on https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
package scfui

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Serve() {
	router = gin.Default()
	router.LoadHTMLGlob("pkg/ui/templates/*")

	initializeRoutes()

	router.Run()
}

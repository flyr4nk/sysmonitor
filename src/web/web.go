package web

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


func Serve() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLGlob("*.html")
	router.GET("/", Server.GetIndexPage)
	router.GET("/info", Server.GetData)
	port := fmt.Sprintf("%s", conf.Basic.ListenPort)
	router.Run(port)
}


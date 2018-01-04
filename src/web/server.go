package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"metrics"
	"confmgr"
)

var conf = confmgr.GetConf()

type server struct {
}

var Server = server{}

func (*server) GetData (c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"System:":   metrics.GetSystemMetricViewData(),
		"Processes":     metrics.GetAllProcessInfo(),
		"CurrentConfig": conf,
	})
}

func (*server) GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"System":   metrics.GetSystemMetricViewData(),
		"Processes":     metrics.GetAllProcessInfo(),
		"CurrentConfig": conf,
	})
}
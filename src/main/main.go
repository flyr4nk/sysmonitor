package main

import (
	"web"
	"jobs"
	"confmgr"
)

var conf = confmgr.GetConf()

func main() {
	jobs.MonitorJob()
	jobs.FileCleanerJob()
	if conf.Basic.EnableWebUi {
		web.Serve()
	} else {
		select {}
	}
}

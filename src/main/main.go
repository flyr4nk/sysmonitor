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
	web.Serve()
}

package jobs

import (
	"github.com/robfig/cron"
	"confmgr"
)

var conf = confmgr.GetConf()

func MonitorJob() {
	if ! conf.WechatConf.Enable {
		return
	}
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		Monitor()
	})
	c.Start()
}

func FileCleanerJob() {
	if ! conf.CleanerConf.Enabled {
		return
	}
	c := cron.New()
	c.AddFunc("@every 1h", func() {
		CleanFile()
	})
	c.Start()
}
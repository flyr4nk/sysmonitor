package jobs

import (
	"net/http"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
	"metrics"
	"confmgr"
)

var hostInfo = confmgr.GetHostInfo()

func httpDo(url, msg string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf("message=%s", msg)))
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	header := conf.WechatConf.Header
	if len(header) > 0 {
		kvs := strings.Split(header, ":")
		if len(kvs) == 2 {
			req.Header.Set(kvs[0], kvs[1])
		}
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(body)
}

func SendWechatMessage(message string) string {
	return httpDo(conf.WechatConf.Url, message)
}

func Monitor() {
	sysInfo := metrics.GetSystemMetric()
	if sysInfo.CpuPercent > conf.SystemConf.CpuLimit {
		SendWechatMessage(fmt.Sprintf("host:%s,ip:%s , System Cpu usage is %.2f too high for the setting limit %.2f",
			hostInfo.HostName, hostInfo.Ipv4Addr, sysInfo.CpuPercent,  conf.SystemConf.CpuLimit))
	}

	if sysInfo.DiskPercent > conf.SystemConf.DiskLimit {
		SendWechatMessage(fmt.Sprintf("host:%s,ip:%s, System Disk usage is %.2f too high for the setting limit %.2f",
			hostInfo.HostName, hostInfo.Ipv4Addr, sysInfo.DiskPercent,  conf.SystemConf.DiskLimit))
	}

	if sysInfo.MemPercent > conf.SystemConf.MemLimit {
		SendWechatMessage(fmt.Sprintf("host:%s, ip:%s , System Memory usage is %.2f too high for the setting limit %.2f",
			hostInfo.HostName, hostInfo.Ipv4Addr, sysInfo.MemPercent,  conf.SystemConf.MemLimit))
	}

	if sysInfo.SystemLoad > conf.SystemConf.SystemLoad {
		SendWechatMessage(fmt.Sprintf("host:%s, ip:%s , System Load is %.2f too high for the setting limit %.2f",
			hostInfo.HostName, hostInfo.Ipv4Addr, sysInfo.SystemLoad,  conf.SystemConf.SystemLoad))
	}
	if len(conf.Process) < 1 {
		return
	}
	for _, v := range conf.Process {
		data := metrics.Monitor.GetProcessInfo(v.Name)
		if ! data.Exists || !data.IsRunning {
			SendWechatMessage(fmt.Sprintf("host:%s, ip:%s, Process name : %s,  the process that you watched is not running",
				hostInfo.HostName, hostInfo.Ipv4Addr, v.Name))
			continue
		}
		if data.OpenFileNum > v.FileLimit {
			SendWechatMessage(fmt.Sprintf("host:%s, ip:%s, Process name: %s, Open file num is %d too many for the setting limit %d",
				hostInfo.HostName, hostInfo.Ipv4Addr, data.Name,  data.OpenFileNum,  v.FileLimit))
		}
		if data.MemPercent > float64(v.MemLimit) {
			SendWechatMessage(fmt.Sprintf("host:%s, ip:%s, Process name: %s, Memory usage is %.2f too high for the setting limit %.2f",
				hostInfo.HostName, hostInfo.Ipv4Addr,  data.Name, data.MemPercent, v.MemLimit))
		}
		if data.CpuPercent > float64(v.CpuLimit) {
			SendWechatMessage(fmt.Sprintf("host %s, ip:%s, Process name: %s, Cpu usage is %.2f too high for the setting limit %.2f",
				hostInfo.HostName, hostInfo.Ipv4Addr, data.Name,  data.CpuPercent,  v.CpuLimit))
		}
		if data.ConnNum > v.ConnLimit {
			SendWechatMessage(fmt.Sprintf("host %s, ip:%s , Process name: %s, Open file num is %d too high for the setting limit %d",
				hostInfo.HostName, hostInfo.Ipv4Addr, data.Name, data.ConnNum,  v.ConnLimit))
		}
		if data.ThreadNum > v.ThreadLimit {
			SendWechatMessage(fmt.Sprintf("host %s, ip:%s ,  Process name: %s,Thread num is %d too high for the setting limit %d",
				hostInfo.HostName, hostInfo.Ipv4Addr, data.Name, data.ThreadNum,  v.ThreadLimit))
		}
	}
}

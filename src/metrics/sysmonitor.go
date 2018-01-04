package metrics

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"confmgr"
)

const (
	B  int64 = 1 << (iota * 10)
	KB
	MB
	GB
)

var conf = confmgr.GetConf()

type SystemMetricViewData struct {
	CpuPercent  string
	MemPercent  string
	DiskPercent string
	SystemLoad	string
}

type SystemMetric struct {
	CpuPercent  float32
	MemPercent  float32
	DiskPercent float32
	SystemLoad	float32
}


type monitorMetric struct {
}

var Monitor = monitorMetric{}

func (*monitorMetric) GetCpuPercent() float64 {
	p, _ := cpu.Percent(0, false)
	return p[0]
}

func (*monitorMetric) GetDiskPercent(path string) float64 {
	dt, _ := disk.Usage(path)
	return float64(dt.Used) / float64(dt.Total) * 100
}

func (*monitorMetric) GetMemPercent() float64 {
	m, _ := mem.VirtualMemory()
	return float64(m.Used) / float64(m.Total) * 100
}

func (*monitorMetric) GetOpenFileNum() int32 {
	return int32(countOpenFiles())
}

func (*monitorMetric) GetSystemLoad() float32 {
	return getSystemLoad()
}

func GetSystemMetricViewData() *SystemMetricViewData {
	m := new(SystemMetricViewData)
	m.CpuPercent = fmt.Sprintf("%.2f%%", Monitor.GetCpuPercent())
	m.DiskPercent = fmt.Sprintf("%.2f%%", Monitor.GetDiskPercent("/"))
	m.MemPercent = fmt.Sprintf("%.2f%%", Monitor.GetMemPercent())
	m.SystemLoad = fmt.Sprintf("%.2f", Monitor.GetSystemLoad())
	return m
}

func GetSystemMetric() *SystemMetric {
	m := new(SystemMetric)
	m.CpuPercent = float32(Monitor.GetCpuPercent())
	m.DiskPercent = float32(Monitor.GetDiskPercent("/"))
	m.MemPercent = float32(Monitor.GetMemPercent())
	m.SystemLoad =  Monitor.GetSystemLoad()
	return m
}

func (*monitorMetric) GetProcessInfo(name string) *ProcessInfo {
	info := new(ProcessInfo)
	pid := getPidByName(name)
	if pid < 1 {
		info.Exists = false
		return info
	}
	info.Exists = true
	p, _ := process.NewProcess(int32(pid))
	isRun, _ := p.IsRunning()
	info.IsRunning = isRun
	memInfo, _ := p.MemoryInfo()
	if memInfo != nil {
		info.Memory = memInfo.RSS
	} else {
		info.Memory = 0
	}
	memPercent, _ := p.MemoryPercent()
	info.Pid = int(p.Pid)
	processName, _ := p.Name()
	info.Name = processName
	username, _ := p.Username()
	info.User = username
	info.MemPercent = float64(memPercent)
	cpuPercent, _ := p.Percent(0)
	info.CpuPercent = cpuPercent
	cnns, _ := p.Connections()
	connNum := len(cnns)
	info.ConnNum = connNum
	fdNum, _ := p.NumFDs()
	info.OpenFileNum = int(fdNum)
	tNum, _ := p.NumThreads()
	info.ThreadNum = int(tNum)
	return info
}

const (
	maxProcessNum = 100
)

func GetAllProcessInfo() []*ProcessInfoView {
	result := make([]*ProcessInfoView, 0, maxProcessNum)
	if len(conf.Process) < 1 {
		return result
	}
	for _, v := range conf.Process {
		data := Monitor.GetProcessInfo(v.Name)
		if ! data.Exists {
			continue
		}
		result = append(result, data.ToViewData())
	}
	return result
}
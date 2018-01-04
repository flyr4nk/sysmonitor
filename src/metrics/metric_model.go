package metrics

import (
	"fmt"
	"encoding/json"
)

type ProcessInfo struct {
	Pid			int
	Name		string
	Memory		uint64
	MemPercent	float64
	User		string
	CpuPercent	float64
	ThreadNum	int
	OpenFileNum	int
	ConnNum		int
	IsRunning	bool
	Exists		bool
}


type ProcessInfoView struct {
	Pid			int
	Name		string
	Memory		string
	MemPercent	string
	User		string
	CpuPercent	string
	ThreadNum	int
	OpenFileNum	int
	ConnNum		int
	IsRunning	bool
	Exists		bool
}

func (info *ProcessInfo) ToViewData() *ProcessInfoView{
	data := new(ProcessInfoView)
	data.Exists		= info.Exists
	if info.Exists {
		data.IsRunning	= info.IsRunning
		data.Name = info.Name
		data.User = info.User
		data.Pid  = info.Pid
		if int64(info.Memory) / GB > 1 {
			data.Memory = fmt.Sprintf("%.2fGB", float64(info.Memory) / float64(GB))
		} else if int64(info.Memory) / MB > 1 {
			data.Memory = fmt.Sprintf("%.2fMB", float64(info.Memory) / float64(MB))
		} else {
			data.Memory = fmt.Sprintf("%.2fKB", float64(info.Memory) / float64(KB))
		}
		data.MemPercent = fmt.Sprintf("%.2f%%", info.MemPercent)
		data.CpuPercent = fmt.Sprintf("%.2f%%", info.CpuPercent)
		data.ThreadNum	= info.ThreadNum
		data.OpenFileNum = info.OpenFileNum
		data.ConnNum	= info.ConnNum
	}
	return data
}

func ( view *ProcessInfoView) String() string {
	result, _ := json.Marshal(view)
	return string(result)
}
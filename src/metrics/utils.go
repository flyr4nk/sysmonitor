package metrics

import (
	"strings"
	"runtime"
	"os/exec"
	"fmt"
	"log"
	"strconv"
	"regexp"
)

func countOpenFiles() int {
	if strings.EqualFold("windows", string(runtime.GOOS)) {
		return 0
	}

	out, err := exec.Command("/bin/sh", "-c", "lsof | wc -l").Output()
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}


func getSystemLoad() float32 {
	if strings.EqualFold("windows", string(runtime.GOOS)) {
		return 0
	}
	out, err := exec.Command("/bin/sh", "-c", "uptime|awk '{print $11}'|cut -d ',' -f1").Output()
	if err != nil {
		log.Printf("Getting system load failed")
		return 0
	}
	loadStr := string(out)
	result, err := strconv.ParseFloat(loadStr[0:len(loadStr) - 1], 32)
	if err != nil {
		log.Printf("error hapeend %v ", err)
		return 0
	}
	return float32(result)
}



var pidRegex = regexp.MustCompile("\\s\\d+\\s")

func getPidByName(name string) int{
	if strings.EqualFold("windows", string(runtime.GOOS)) {
		out, err := exec.Command("cmd", "/c", fmt.Sprintf("tasklist | findstr %s", name)).Output()
		if err != nil {
			log.Printf("get pid failed: %v\n", err)
		}
		processRawStr := string(out)
		if len(processRawStr) < 1 {
			return 0
		}
		resultStr := pidRegex.FindAllString(processRawStr, -1)
		if len(resultStr) < 1 {
			return 0
		}
		pidStr := resultStr[0]
		pid, err :=  strconv.Atoi(strings.TrimSpace(pidStr))
		if err != nil {
			log.Printf("Get pid error: %v\n", err)
		}
		return pid
	} else {

		// ps -ef | grep "nova" | grep -v grep | awk '{print $2}'
		out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("ps -ef | grep \"%s\" | grep -v grep | awk '{print $2}'", name)).Output()
		if err != nil {
			log.Printf("get pid failed: %v\n", err)
		}
		rawPidStr := strings.TrimSpace(string(out))
		if len(rawPidStr) < 1 {
			return 0
		}
		pidStrs := strings.Split(rawPidStr, "\n")
		if len(pidStrs) < 1{
			return 0
		}
		pid , err := strconv.Atoi(pidStrs[len(pidStrs) - 1])
		if err != nil {
			log.Printf("get pid failed: %v\n", err)
		}
		return pid
	}
	return 0
}
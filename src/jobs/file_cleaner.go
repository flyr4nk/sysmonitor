package jobs

import (
	"strings"
	"strconv"
	"log"
	"path/filepath"
	"os"
	"time"
)

const fileNum = 100

func CleanFile() {
	tStr := conf.CleanerConf.Time
	// 如果存在日期配置的字符串
	var t int
	var err error
	if strings.Index(tStr, "d")+strings.Index(tStr, "h") > -1 {
		ts := tStr[0:len(tStr)-1]
		t, err = strconv.Atoi(ts)
		if err != nil {
			log.Printf("Wrong format:%v\n", err)
			return
		}
		if strings.Contains(tStr, "d") {
			t = t * 24 * 3600
		} else {
			t = t * 3600
		}
	} else {
		return
	}
	dir := conf.CleanerConf.FileDir
	files := ListDir(dir, "")
	for _, v := range files {
		f, err := os.Open(v)
		if err != nil {
			log.Printf("error reading file : %s, %v\n", v, err)
		}
		fstat, err := f.Stat()
		if err != nil {
			log.Printf("error getting file info : %s, %v\n", v, err)
		}
		// 如果时间超过了规定的时间, 则删掉
		if time.Now().UnixNano() - fstat.ModTime().UnixNano() > int64(time.Second) * int64(t) {
			log.Printf("File cleaner: deleting file %s", v)
			os.Remove(v)
		}
	}
}

func ListDir(dirPth, suffix string) []string {
	files := make([]string, 0, fileNum)
	suffix = strings.ToUpper(suffix)
	err := filepath.Walk(dirPth,
		func(filename string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
				return nil
			}
			if len(suffix) > 0 {
				if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
					files = append(files, filename)
				}
			} else {
				files = append(files, filename)
			}
			return nil
	})
	if err != nil {
		return make([]string, 0, 0)
	}
	return files
}

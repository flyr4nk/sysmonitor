package confmgr

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)


var conf *Config

var once sync.Once

func readConfFromFile() {
	if nil == conf {
		conf = &Config{}
		fileContent, err := ioutil.ReadFile("conf.yaml")
		if nil != err {
			log.Printf("read file error! #%v ", err)
			return
		}
		err = yaml.Unmarshal(fileContent, conf)
		if nil != err {
			log.Fatal("Read yaml error!:", err)
		}
	}
}

func GetConf() *Config {
	if nil == conf {
		readConfFromFile()
	}
	return conf
}


func (c *Config) String() string {
	content, err :=  yaml.Marshal(c)
	if nil != err {
		log.Printf("to yaml string failed: %v", err)
	}
	return string(content)
}
package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	version = kingpin.Flag("version", "software version").Default("0.1.0").String()
)

func init() {
	kingpin.Parse()

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("Error occurred: %s", r))
		}
	}()
	fmt.Println(*version)

	var s interface{}
	s = "xxx"
	_, ok := s.(string)
	if ok {
		fmt.Println("OK")
	}

	//var arry interface{}

	array := [...]string{}

	//_, ok = array.([...]string)
	if ok {
		for _, a := range array {
			fmt.Println(a)
		}
	}

	// Read Configuration from File
	file, err := ioutil.ReadFile("config.yml")
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Errorf("Load configuration failed: %s", err)
		os.Exit(0)
	}
	log.Infoln(Config)

}

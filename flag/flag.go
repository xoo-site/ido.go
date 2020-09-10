package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
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

}

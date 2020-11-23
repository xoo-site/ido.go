package main

import "fmt"

func main() {
	//var s = "  xxx\n"
	//fmt.Println(strings.TrimSpace(s))
	var diskStats map[string]map[int]string
	diskStats = map[string]map[int]string{}
	diskStats["sda"] = map[int]string{
		1: "xxx",
		2: "xxx",
		3: "xxx",
	}
	diskStats["sdb"] = map[int]string{
		1: "cc",
		2: "cc",
		3: "cc",
		4: "cc",
		5: "cc",
	}
	diskStats["sdc"] = map[int]string{
		0: "c",
	}

	for dev, _ := range diskStats {
		max := 20
		cols := len(diskStats[dev])
		for i := 0; i < max-cols; i++ {
			diskStats[dev][cols+i] = "0"
		}
	}

	fmt.Println(diskStats)

}

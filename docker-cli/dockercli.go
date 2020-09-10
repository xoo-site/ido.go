package main

import "fmt"

func init() {
	fmt.Println("Init called")
}

func main() {
	// 获取docker状态
	//ctx := context.Background()
	//cli, err := client.NewClient("192.168.1.77", "", )
	//if err != nil {
	//	panic("connect failed")
	//}
	//stats, err := cli.ContainerStats(ctx, "xxx", true)
	//fmt.Println(stats)
	fmt.Println("Main called")
}

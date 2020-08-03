package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

var (
	smtp_account  = "alert.it@qtechglobal.com"
	smtp_password = "xxxhaofan@2008"
	smtp_host     = "smtphm.qiye.163.com:25"
	receivers     = []string{"jianxin.lu@woqutech.com"}
)

func main() {
	auth := smtp.PlainAuth("xx", smtp_account, smtp_password, strings.Split(smtp_host, ":")[0])
	fmt.Println(auth)
	//strings.Split(smtp_host, ":")[0]
	//auth.Start(&smtp.ServerInfo{})
	err := smtp.SendMail(smtp_host, auth, smtp_account, receivers, []byte("Mail from Jeeyshe"))
	fmt.Println(err)
}

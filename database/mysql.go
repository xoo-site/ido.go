package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	name   string
	email  sql.NullString
	mobile sql.NullString
}

var connString = "pig:p7tiULiN0xSp2S03ZHJmHoVBaEYg3NYoRF0h4O7TIEk=@tcp(192.168.1.77:9307)/qdata_cloud"

func main() {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println(err)
		return
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	rows, err := db.QueryContext(timeout, "select name, email, mobile from auth_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.name, &user.email, &user.mobile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user)
	}
}

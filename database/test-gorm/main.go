package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type BaseModel struct {
	CreateAt time.Time `gorm:"default:"` // 记录创建时间
	UpdateAt time.Time // 记录最后修改时间
}

type AuthUser struct {
	ID       uint `gorm:"default:uuid_generate_v3()"`
	Username string
	Password string
	NickName string `gorm:"default:galeone"`
	BaseModel
}

type Article struct {
	ID      uint
	Title   string
	Context string
	BaseModel
}

func init() {

}

func main() {
	// 创建连接引擎
	db, err := gorm.Open(
		sqlite.Open("db.sqlite3"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// 设置自动迁移
	_ = db.AutoMigrate(&AuthUser{})
	_ = db.Migrator().CreateTable(&AuthUser{})

	// 插入单条记录 db.Create
	db.Create(&AuthUser{
		Username: "Jeyrce.Lu",
		Password: "cljslrl0620",
		NickName: "Mr.Lu",
	})
	// 批量插入记录
	db.Create(&[]AuthUser{
		{Username: "shangchenhui", Password: "12345"},
		{Username: "yanqing", Password: "12345"},
		{Username: "yutingting", Password: "12345"},
	})

	// 使用map创建 db.Model.Create
	db.Model(&AuthUser{}).Create(&[]map[string]interface{}{
		{"Username": "huanglinshan", "Password": "1234567"},
		{"Username": "fanwaowen", "Password": "1234567"},
		{"Username": "sunchaoyang", "Password": "1234567"},
	})

	// 批量写入
	db.CreateInBatches(
		[]AuthUser{
			{Username: "woqutech", Password: "woqutech"},
			{Username: "dell", Password: "dell"},
			{Username: "h3c", Password: "h3c"},
		},
		3,
	)
}

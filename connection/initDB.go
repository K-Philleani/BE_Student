package connection

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
const dsn string = "root:123456@(124.70.71.78:3306)/Stitches"

func init() {
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	log.Println(">>>>>>>>>>>>>>>>>>>数据库已连接.......<<<<<<<<<<<<<<<<<<<<")
}

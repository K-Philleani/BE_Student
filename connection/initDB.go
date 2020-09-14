package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	dsn := "root:00000000@tcp(127.0.0.1:3306)/Student"
	SqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("数据库连接失败， err: ", err)
		return
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Println("数据库连接失败, err: ", err)
		return
	}
}

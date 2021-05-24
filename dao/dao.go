package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"trainingcamp/pkg/setting"
)

type Dao struct {
	Id int
	CreatedOn int
	CreatedBy string
	ModifiedOn int
	ModifiedBy string
	DeletedOn int
}

var db *sql.DB

func Setup() {
	var err error

	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		setting.Config.Mysql.User,
		setting.Config.Mysql.Password,
		setting.Config.Mysql.Host,
		setting.Config.Mysql.Port,
		setting.Config.Mysql.DbName,
		))
	if err != nil {
		log.Fatalf("%v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
}

func Close() {
	db.Close()
}

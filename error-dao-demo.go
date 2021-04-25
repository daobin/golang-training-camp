package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TestData struct {
	Name string `db:"name"`
	Age int `db:"age"`
}

var Db *sqlx.DB

func init() {
	Db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.21:3306)/test_data")
	if err != nil {
		fmt.Println("Connect Mysql Failed, ", err)
		return
	}

	defer Db.Close()
}

func main()  {
	var td []TestData
	err := Db.Get(td, "select `name`, `age` from `test_data` where `age` = ?", 18)
	if err != nil {
		fmt.Println("Exec Failed: ", err)
	}

	fmt.Println(td)
}


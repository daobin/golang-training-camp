package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Test struct {
	id int
	name string
	age int
}

var db *sql.DB

func initDB() (err error) {
	db, err := sql.Open("mysql", "xxx:xxx@tcp(127.0.0.1:3306)/kratos")
	if err != nil {
		// 直接返回错误给调用者
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		// 直接返回错误给调用者
		return
	}

	return nil
}

func queryRow(id int) (Test, error) {
	var test Test
	sql := "select `id`, `name`, `age` from `test` where `id` = ?"
	row := db.QueryRow(sql, id)
	err := row.Scan(&test.id, &test.name, &test.age)

	// 当查询无结果错误时，直接返回给调用者判断。想像不到包装错误的使用场景，请老师解惑。
	return test, err
}

func main()  {
	err := initDB()
	if err != nil {
		fmt.Printf("DB init failed, err >> %v\n", err)
		return
	}

	// 查询数据居然被360识别为木马，请问老师，在 win 系统上开发如何解决

	//test, err := queryRow(1)
	//if err != nil {
	//	fmt.Printf("DB query failed, err >> %v\n", err)
	//	return
	//}
	//
	//fmt.Println(test)
}


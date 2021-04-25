package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	name string
	age int
}

var db *sql.DB

func initDB() (err error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
		return
	}

	defer db.Close()

	return nil
}

func queryRow(id int) (User, error) {
	var user User
	sql := "select `id`, `name`, `age` from `user` where `id` = ?"
	row := db.QueryRow(sql, id)
	err := row.Scan(&user.id, &user.name, &user.age)

	return user, err
}

func main()  {
	err := initDB()
	if err != nil {
		fmt.Println("DB init failed, err = %v\n", err)
		return
	}

	user, err := queryRow(1)
	if err != nil {
		fmt.Println("DB query failed, err = %v\n", err)
		return
	}

	fmt.Println(user)
}


package Initializers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sample")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("DB is connected Successfully")
	return db
}

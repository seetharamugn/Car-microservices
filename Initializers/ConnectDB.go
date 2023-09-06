package Initializers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDb() *sql.DB {

	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("DB is connected Successfully")
	return db
}

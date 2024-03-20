package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func Init() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dailyinsights")
	if err != nil {
		log.Fatal("Error while creating the connection : ", err)
	}
	Conn = db
	fmt.Println("Database Connection Successfull.")
	AutoMigration()
}

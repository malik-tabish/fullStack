package db

import (
	"log"
)

func AutoMigration() {
	users()
}

func users() {
	query := "CREATE TABLE IF NOT EXISTS users(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100), username VARCHAR(100) UNIQUE, password VARCHAR(16))"

	_, err := Conn.Query(query)
	if err != nil {
		log.Fatal("Error while creating the user table : ", err)
	}
	log.Println("User table created")
}

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "qwe123"
	NETWORK  = "tcp"
	SERVER   = "1207.0.0.1"
	PORTS    = 3306
	DATABASE = "test"
)

func SqlConn() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORTS, DATABASE)
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Print("connected successful")
	return db
}

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
	DATABASE = "threekingdoms"
)

func SqlConn() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORTS, DATABASE)
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("connected successful")
	res, _ := db.Exec("SELECT * FROM `roles`")
	fmt.Println(res)
	return db
}

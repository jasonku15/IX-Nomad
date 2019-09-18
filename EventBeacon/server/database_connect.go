package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

var db *sql.DB

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config[dbuser], config[dbpass], config[dbhost], config[dbport], config[dbname])

	db, err = sql.Open("mysql", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[dbhost] = "localhost"
	conf[dbport] = "3306"
	conf[dbuser] = "root"
	conf[dbpass] = ""
	conf[dbname] = "go-mysql-crud"
	return conf
}

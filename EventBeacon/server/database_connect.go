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
	// host, ok := os.LookupEnv(dbhost)
	// if !ok {
	// 	panic("DBHOST environment variable required but not set")
	// }
	// port, ok := os.LookupEnv(dbport)
	// if !ok {
	// 	panic("DBPORT environment variable required but not set")
	// }
	// user, ok := os.LookupEnv(dbuser)
	// if !ok {
	// 	panic("DBUSER environment variable required but not set")
	// }
	// password, ok := os.LookupEnv(dbpass)
	// if !ok {
	// 	panic("DBPASS environment variable required but not set")
	// }
	// name, ok := os.LookupEnv(dbname)
	// if !ok {
	// 	panic("DBNAME environment variable required but not set")
	// }
	// conf[dbhost] = host
	// conf[dbport] = port
	// conf[dbuser] = user
	// conf[dbpass] = password
	// conf[dbname] = name

	conf[dbhost] = "localhost"
	conf[dbport] = "3306"
	conf[dbuser] = "root"
	conf[dbpass] = ""
	conf[dbname] = "go-mysql-crud"
	return conf
}

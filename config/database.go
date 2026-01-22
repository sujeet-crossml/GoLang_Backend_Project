package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB
func ConnectDB(){
	dsn := "root:password@tcp(localhost:3306)/shop?parseTime=true"
	db, err := sql.Open("mysql", dsn)	
	if err != nil{
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}
	DB = db

}
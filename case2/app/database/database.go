package database

import (
	"case2/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(creds config.DBCredential) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", creds.UserName, creds.Password, creds.Host, creds.Port, creds.DBName))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(15 * time.Minute)

	return db
}

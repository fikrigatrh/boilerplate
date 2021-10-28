package db

import (
	"boilerplate/config/env"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var SqlDB *sql.DB

func init() {
	var err error
	var configDb = env.Config.PostgresConfig

	if SqlDB != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovery error")
		}
	}()

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=UTC", configDb.User, configDb.Password, configDb.Host, configDb.Port, configDb.Name, "disable",
	)


	SqlDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	SqlDB.SetMaxOpenConns(50)
	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetConnMaxLifetime(time.Duration(20) * time.Second)

	log.Print("success connect to db with sql lib")

}

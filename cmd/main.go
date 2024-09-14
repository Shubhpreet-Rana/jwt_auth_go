package main

import (
	"database/sql"
	"log"

	"github.com/Shubhpreet-Rana/jwt_auth_go/cmd/api"
	"github.com/Shubhpreet-Rana/jwt_auth_go/config"
	"github.com/Shubhpreet-Rana/jwt_auth_go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	envConfig := config.Env

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 envConfig.DBUser,
		Passwd:               envConfig.DBPassword,
		Addr:                 envConfig.DBAddress,
		DBName:               envConfig.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	initStorageDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorageDatabase(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}

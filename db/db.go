package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewMySqlDB() (*sql.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/go_test_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Database connection failed:", err)
		return nil, err
	}

	DB = db
	fmt.Println(" Database connection successful.")
	return db, nil
}

/*
 Go Module Create Command:
	go mod init your-module-name

 Gin Framework Install Command:
	go get -u github.com/gin-gonic/gin

 MySQL Driver Install Command:
	go get github.com/go-sql-driver/mysql
*/

/*
 Connection String (DSN) Format:
	username:password@tcp(host:port)/database

 Example:
	1. Username:      root
	2. Password:      secret
	3. Protocol:      tcp
	4. Host & Port:   127.0.0.1:3306
	5. Database Name: your_db_name

 Final DSN:
	root:secret@tcp(127.0.0.1:3306)/your_db_name
*/

/*

1.db.go
2. user_repo.go
3.user.go
4.auth_service.go
5. util.go
7.auth_handele
8. main.go



*/

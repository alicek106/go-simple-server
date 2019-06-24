package mysql_conn

import (
	"database/sql" 
    _ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func CreateMysqlConn() *sql.DB {
	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(13.125.242.14:3306)/mydb")
	if err != nil{
		log.Fatal(err)
	}

	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}

	return db
}
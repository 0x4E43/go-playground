package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func Db() {

	fmt.Println(os.Getenv("DB_NAME"))
	user := "nimai" //need to be taken from .env
	passw := ""
	db, err := sql.Open("mysql", user+":"+passw+"@tcp(10.10.10.11:3306)/report_generator")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second * 30)

	query := "SELECT * FROM node_master"
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}

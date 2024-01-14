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

	DB_NAME := os.Getenv("DB_NAME")
	DB_URL := os.Getenv("DB_URL")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_URL+":"+DB_PORT+")/"+DB_NAME)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second * 30)

	query := "SELECT * FROM node_master"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(rows.Columns()) //prints db columns
	// needs to scan results

	var dbresult []NodeMaster

	for rows.Next() {
		var nodeData NodeMaster
		if err := rows.Scan(&nodeData.Id, &nodeData.DeviceId, &nodeData.ControllerId,
			&nodeData.IEEAddress, &nodeData.Intensity, &nodeData.Occupancy,
			&nodeData.IsIntensitySent, &nodeData.IsOccupancySent); err != nil {
			panic(err)
		}
		dbresult = append(dbresult, nodeData)
	}
	fmt.Println(dbresult)
}

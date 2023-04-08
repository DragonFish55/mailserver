package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createDbConnect(dbname string) (db *sql.DB, err_out error) {
	sqlconn := "host=" + dbaddr + " port=" + dbport + " user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " sslmode=disable"
	fmt.Println(sqlconn)
	dbmain, err := sql.Open(dbtype, sqlconn)
	if err != nil {
		fmt.Println("eror")
	}
	err = dbmain.Ping()
	return dbmain, err
}

//func register_user()

func checkUserValid(password string, key string) bool {
	return true

}

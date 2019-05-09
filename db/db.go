package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Conn() *sql.DB {
	if db != nil {
		return db
	}
	var err error
	url := "postgres://bkadffuj:zP2xt2JXZIH_O4Oua3PsKcHRiIlfjcCP@baasu.db.elephantsql.com:5432/bkadffuj"
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can't connect database : ", err)
	}
	return db
}

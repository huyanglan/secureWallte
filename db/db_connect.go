package db

import (
	"database/sql"
	"log"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

// db
const (
	userName = "root"
	passwd   = "Aka@gp!1231"
	ip       = "192.168.31.245"
	port	 = "3306"
	maxConnectNum = 2000
	maxIdealConnectNum = 500
)
var DB *sql.DB
// init mysql db
func InitDB(dbName string) *sql.DB{
	url := []string{userName,":",passwd,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"}
	path := strings.Join(url, "")
	DB, _ = sql.Open("mysql", path)
	// set max connect number
	DB.SetConnMaxLifetime(maxConnectNum)
	// set max
	DB.SetMaxIdleConns(maxIdealConnectNum)

	if err := DB.Ping(); err != nil {
		log.Panic("Connect mysql fail")
	}
	return DB
}

func Insert(sql string, ) {

}




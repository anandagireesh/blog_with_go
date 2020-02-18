package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB


//DbConnection -
func  DbConnection(){
	var err error
	Db, err = sql.Open("mysql","14nqZ8a7Xc:rSxe73oDwr@tcp(remotemysql.com:3306)/14nqZ8a7Xc")

	if err != nil {
		panic(err.Error())
	}

	//fmt.Println("OK")

}

//GetConnection -
func GetConnection() *sql.DB{
	return Db
}

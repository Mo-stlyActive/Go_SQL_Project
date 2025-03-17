package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// have to connect my own sql table here
func Connect() {
	//Make sure to replace username, password, localhost, 1433, and dbname with your actual SQL Server credentials and database information.
	d, err := gorm.Open("mysql", "sqlserver://username:password@localhost:1433?database=dbname")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

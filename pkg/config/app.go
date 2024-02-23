package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// creating a coonect function wihich help us with the making connection with the database

func Connect() {
	d, err := gorm.Open("mysql", "root:1234567890@tcp(localhost:3306)/simpleinterset?charset=utf8&parseTime=True&loc=Local")
	// here username is root and password then we have provided the port where server is running
	// for that we have to first create database name simpleinterset then coonect it to mysql server
	if err != nil {
		panic(err)
	}
	db = d
	// now db is assigned new connection

}
func GetDB() *gorm.DB {
	return db
}

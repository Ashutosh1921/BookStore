package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// creating a coonect function wihich help us with the making connection with the database

func Connect() {
	d, err := gorm.Open("mysql", "ashutosh0010:12345@123/simpleinterset?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
	// now db is assigned new connection

}
func GetDB() *gorm.DB {
	return db
}

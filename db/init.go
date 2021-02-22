package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db  *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func ConnectDb() {

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))

	d, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("I am connected")
		db = d
	}

}

package config

import (
	"github.com/jinzhu/gorm"                  //which is ORM
	_ "github.com/jinzhu/gorm/dialects/mysql" //for mysql
)

var (

	//which will interact with db
	db *gorm.DB
)

func Connect() {
	//help us to connection with mysql database
	d, err := gorm.Open("mysql", "aniket:Qwerty@123/mydb?charset=utf8&parseTime=True&loc=Local") //which is used to open my connect with DB. simplerest is a table in db
	if err != nil {
		panic(err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}

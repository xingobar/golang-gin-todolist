package models

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:@/gin_todo?parseTime=true")
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
}

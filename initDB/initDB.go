package initDB

import (
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:@/gin_todo")
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
}

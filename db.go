package ngdc

import (
	"log"

	"gorm.io/gorm"
)

//	type User struct{
//		Name string
//	 	Age int
//	  }

//dsn = database source name or yoi can call DB_url

//	dialector := postgres.Open(dsn)
//	db := gdc.ConnectDB(dialector, &User{}) {*gorm.DB}

var db *gorm.DB

func ConnectDB(dialector gorm.Dialector, model ...interface{}) (*gorm.DB, error) {

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panic("Database connection Failed!")
	}
	log.Println("Database Connected")
	db.AutoMigrate(model...)
	return db, err
}

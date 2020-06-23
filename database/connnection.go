package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// GORM MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectToDB open connection to database
func ConnectToDB() *gorm.DB {
	db, err := gorm.Open("mysql", viper.Get("dbUser").(string)+":"+viper.Get("dbPassword").(string)+"@tcp("+viper.Get("dbHost").(string)+")/"+viper.Get("dbName").(string)+"?charset=utf8mb4&parseTime=True&loc=Local")

	// unable to connect
	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	err = db.DB().Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

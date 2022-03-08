package database

import (
	"log"
	"rest-api-mysql/model"
	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connect to database successfully!")
	return nil
}

//Migrate create/update database table
func Migrate(table *model.User){
	Connector.AutoMigrate(&table)
	log.Println("Table migrated!")
}
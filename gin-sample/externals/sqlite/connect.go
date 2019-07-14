package sqlite

import (
	"til-golang_learning/gin-sample/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if !db.HasTable(&gateway.User{}) {
		// Migrate the schema
		db.AutoMigrate(&gateway.User{})
		// Create
		db.Create(&gateway.User{Username: "kuzu", Password: "kuzukuzu"})
	}

	return db
}

func CloseConn() {
	db.Close()
}

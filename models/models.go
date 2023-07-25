package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

type User struct {
	gorm.Model
	Order []Order
	Data  string `sql:"type:JSONB NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	User User
	Data string `sql:"type:JSON NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM create table tables with plural names. Use this to suppress it
func (User) TableName() string {
	return "user"
}

func (Order) TableName() string {
	return "order"
}

func IniDB() (*gorm.DB, error) {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	} else {
		/*
			//The below AutoMigrate is quivalent to this
			if !db.HasTable("user") {
				db.CreateTable(&User{})
			}

			if !db.HasTable("order") {
				db.CreateTable(&Order{})
			}
		*/
		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}
}
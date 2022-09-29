package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initial() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	if err := sqlDB.Ping(); err != nil {
		panic("failed to ping database")
	}
	return db
}

// func (db *gorm.DB) runMigrate() {

// }

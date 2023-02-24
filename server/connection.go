package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect Connection to the Sqlite database
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("xplane11-webRTC.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

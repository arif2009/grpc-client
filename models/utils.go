package models

import (
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// Include this to access secrets (which should contain your DB password)
	// "github.com/synspective/syns-platform-backend-sample-rest/pkg/core"
)

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

var DB *gorm.DB

// Connect to database
func Connect() (err error) {
	// 1. For the purpose of this example, we'll be using SQLite.
	//    - Replace the driver with the appropriate database used.
	//    - If you're using database other than the ones supported by Gorm, change this
	//      to whatever can be used to interact with your database.
	// 2. At this point you can access secrets by importing `core` and using `core.Config`
	if DB, err = gorm.Open(sqlite.Open("/tmp/gorm.db"), &gorm.Config{}); err != nil {
		return errors.New("error connecting to DB")
	}

	// Since we're using gorm, just run auto-migrate each time.
	// It won't do any harm if there are no migrations to run.
	if err = DB.AutoMigrate(
		&User{},
	); err != nil {
		return
	}

	return nil
}

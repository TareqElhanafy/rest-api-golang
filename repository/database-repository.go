package repository

import (
	"RestApi/entity"

	"github.com/jinzhu/gorm"
)

//DatabaseRepository composing interface
type DatabaseRepository interface {
	VideoRepository
	UserRepository
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

// NewDatabaseRepository function connection and migration to the database
func NewDatabaseRepository() DatabaseRepository {
	db, err := gorm.Open("mysql", "root:password@/testapi")
	if err != nil {
		panic("failed to connect to database")
	} else {
		db.AutoMigrate(&entity.Video{}, &entity.Person{})

	}

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close database")
	}
}

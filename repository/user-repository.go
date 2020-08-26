package repository

import (
	"RestApi/entity"
)

//UserRepository interface
type UserRepository interface {
	StoreUser(user entity.Person)
	FindAllUsers() []entity.Person
}

func (db *database) StoreUser(user entity.Person) {
	db.connection.Create(&user)
}
func (db *database) FindAllUsers() []entity.Person {
	var users []entity.Person
	db.connection.Find(&users)
	return users
}

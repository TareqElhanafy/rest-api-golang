package service

import (
	"RestApi/entity"
	"RestApi/repository"
)

//UserService interface
type UserService interface {
	StoreUser(user entity.Person) entity.Person
	FindAllUsers() []entity.Person
}

type userService struct {
	repository repository.UserRepository
}

// NewUserService construction function
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) StoreUser(user entity.Person) entity.Person {
	service.repository.StoreUser(user)
	return user
}
func (service *userService) FindAllUsers() []entity.Person {
	return service.repository.FindAllUsers()
}

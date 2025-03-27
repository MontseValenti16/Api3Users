package repositories

import (
	"errors"
	"API3/registro/src/domain/entities"
)

var users = make(map[int]entities.User)
var nextID = 1

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(user entities.User) (entities.User, error) {
	user.ID = nextID
	nextID++
	users[user.ID] = user
	return user, nil
}

func (repo *UserRepositoryImpl) GetByID(id int) (entities.User, error) {
	user, exists := users[id]
	if !exists {
		return entities.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetAll() ([]entities.User, error) {
	var list []entities.User
	for _, u := range users {
		list = append(list, u)
	}
	return list, nil
}

func (repo *UserRepositoryImpl) Update(user entities.User) (entities.User, error) {
	_, exists := users[user.ID]
	if !exists {
		return entities.User{}, errors.New("usuario no encontrado")
	}
	users[user.ID] = user
	return user, nil
}

func (repo *UserRepositoryImpl) Delete(id int) error {
	_, exists := users[id]
	if !exists {
		return errors.New("usuario no encontrado")
	}
	delete(users, id)
	return nil
}

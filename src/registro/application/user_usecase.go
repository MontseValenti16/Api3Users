package user_usecase

import (
	"API3/src/registro/domain/entities"
	"API3/src/registro/domain/repositories"
)

type UserUseCase struct {
	Repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) RegisterUser(user entities.User) (entities.User, error) {
	return uc.Repo.Create(user)
}

func (uc *UserUseCase) GetUser(id int) (entities.User, error) {
	return uc.Repo.GetByID(id)
}

func (uc *UserUseCase) GetUsers() ([]entities.User, error) {
	return uc.Repo.GetAll()
}

func (uc *UserUseCase) UpdateUser(user entities.User) (entities.User, error) {
	return uc.Repo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id int) error {
	return uc.Repo.Delete(id)
}

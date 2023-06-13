package usecase

import (
	"koizumi55555/go-restapi/src/entitiy"
)

type UserUsecase struct {
	PostgresIf PostgresIf
}

func NewUserUsecase(pif PostgresIf) UserUsecase {
	return UserUsecase{
		PostgresIf: pif,
	}
}

func (uuc UserUsecase) GetUser(id string) (user entitiy.User, err error) {
	return uuc.PostgresIf.GetUserDB(id)
}

func (uuc UserUsecase) DeleteUser(id string) (err error) {
	return uuc.PostgresIf.DeleteUserDB(id)
}

func (uuc UserUsecase) UpdateUser(updateUser entitiy.User) (user entitiy.User, err error) {
	return uuc.PostgresIf.UpdateUserDB(updateUser)
}

func (uuc UserUsecase) CreateUser(createUser entitiy.User) (user entitiy.User, err error) {
	return uuc.PostgresIf.CreateUserDB(createUser)
}

func (uuc UserUsecase) GetListUsers() (user []entitiy.User, err error) {
	return uuc.PostgresIf.GetListUsersDB()
}

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

	getUser, err := uuc.PostgresIf.GetUserDB(id)
	if err != nil {
		return err
	}

	return uuc.PostgresIf.DeleteUserDB(getUser.ID)
}

func (uuc UserUsecase) UpdateUser(updateUser entitiy.User) (user entitiy.User, err error) {

	getUser, err := uuc.PostgresIf.GetUserDB(updateUser.ID)
	if err != nil {
		return getUser, err
	}

	return uuc.PostgresIf.UpdateUserDB(getUser, updateUser)
}

func (uuc UserUsecase) CreateUser(createUser entitiy.User) (user entitiy.User, err error) {
	return uuc.PostgresIf.CreateUserDB(createUser)
}

func (uuc UserUsecase) ListUsers() (user []entitiy.User, err error) {
	return uuc.PostgresIf.ListUsersDB()
}

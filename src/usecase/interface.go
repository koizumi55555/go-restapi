package usecase

import "koizumi55555/go-restapi/src/entitiy"

type PostgresIf interface {
	GetUserDB(id string) (user entitiy.User, err error)
	DeleteUserDB(id string) (err error)
	UpdateUserDB(updateUser entitiy.User) (user entitiy.User, err error)
	CreateUserDB(createUser entitiy.User) (user entitiy.User, err error)
	GetListUsersDB() (user []entitiy.User, err error)
}

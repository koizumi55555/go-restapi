package postgres

import (
	"errors"
	"koizumi55555/go-restapi/src/entitiy"
	"strconv"

	"gorm.io/gorm"
)

// ユーザ情報の取得
func (p Postgres) GetUserDB(id string) (user entitiy.User, err error) {

	resultUser := Users{}
	getResult := p.Conn.Find(&resultUser, "id = ?", id)

	if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, getResult.Error
	}

	return entitiy.User{
		ID:    resultUser.ID,
		Name:  resultUser.UserName,
		Email: resultUser.Email,
	}, nil
}

// ユーザの削除
func (p Postgres) DeleteUserDB(id string) (err error) {
	deleteResult := p.Conn.Where("id = ?", id).Delete(&entitiy.User{})

	if errors.Is(deleteResult.Error, gorm.ErrRecordNotFound) {
		return deleteResult.Error
	}
	return nil
}

// ユーザの更新
func (p Postgres) UpdateUserDB(updateUser entitiy.User) (user entitiy.User, err error) {

	resultUser := Users{}
	getResult := p.Conn.Find(&resultUser, "id = ?", updateUser.ID)

	if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, getResult.Error
	}

	resultUser.UserName = updateUser.Name
	resultUser.Email = updateUser.Email
	updateResult := p.Conn.Save(&resultUser)

	if errors.Is(updateResult.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, updateResult.Error
	}

	return entitiy.User{
		ID:    resultUser.ID,
		Name:  resultUser.UserName,
		Email: resultUser.Email,
	}, nil
}

// ユーザの登録
func (p Postgres) CreateUserDB(createUser entitiy.User) (user entitiy.User, err error) {

	resultUser := Users{}
	getResult := p.Conn.Order("id desc").First(&resultUser)

	if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, getResult.Error
	}

	i, _ := strconv.Atoi(resultUser.ID)

	newUser := Users{
		ID:       strconv.Itoa(i + 1),
		UserName: createUser.Name,
		Email:    createUser.Email,
	}

	createResult := p.Conn.Save(&newUser)

	if errors.Is(createResult.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, createResult.Error
	}

	return entitiy.User{
		ID:    newUser.ID,
		Name:  newUser.UserName,
		Email: newUser.Email,
	}, nil
}

// ユーザ一覧取得
func (p Postgres) ListUsersDB() (user []entitiy.User, err error) {
	resultUser := []Users{}
	getResult := p.Conn.Order("id asc").Find(&resultUser)

	if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
		return nil, getResult.Error
	}

	for i := 0; i < len(resultUser); i++ {
		user = append(user, entitiy.User{
			ID:    resultUser[i].ID,
			Name:  resultUser[i].UserName,
			Email: resultUser[i].Email,
		})
	}

	return user, nil
}

package postgres

import (
	"errors"
	"koizumi55555/go-restapi/src/entitiy"
	"strconv"

	"gorm.io/gorm"
)

// ユーザ情報の取得
func (p Postgres) GetUserDB(id string) (user entitiy.User, err error) {

	user1 := Users{}
	result := p.Conn.Find(&user1, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, result.Error
	}

	return entitiy.User{
		ID:    user1.ID,
		Name:  user1.UserName,
		Email: user1.Email,
	}, nil
}

// ユーザの削除
func (p Postgres) DeleteUserDB(id string) (err error) {
	result := p.Conn.Where("id = ?", id).Delete(&entitiy.User{})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

// ユーザの更新
func (p Postgres) UpdateUserDB(updateUser entitiy.User) (user entitiy.User, err error) {

	user1 := Users{}
	result1 := p.Conn.Find(&user1, "id = ?", updateUser.ID)

	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, result1.Error
	}

	user1.UserName = updateUser.Name
	user1.Email = updateUser.Email
	result2 := p.Conn.Save(&user1)

	if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, result2.Error
	}

	return entitiy.User{
		ID:    user1.ID,
		Name:  user1.UserName,
		Email: user1.Email,
	}, nil
}

// ユーザの登録
func (p Postgres) CreateUserDB(createUser entitiy.User) (user entitiy.User, err error) {

	user1 := Users{}
	result1 := p.Conn.Order("id desc").First(&user1)

	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, result1.Error
	}

	user2 := Users{}
	i, _ := strconv.Atoi(user1.ID)
	user2.ID = strconv.Itoa(i + 1)
	user2.UserName = createUser.Name
	user2.Email = createUser.Email
	result2 := p.Conn.Save(&user2)

	if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		return entitiy.User{}, result2.Error
	}

	return entitiy.User{
		ID:    user2.ID,
		Name:  user2.UserName,
		Email: user2.Email,
	}, nil
}

// ユーザ一覧取得
func (p Postgres) GetListUsersDB() (user []entitiy.User, err error) {
	user1 := []Users{}
	result := p.Conn.Order("id asc").Find(&user1)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	for i := 0; i < len(user1); i++ {
		user = append(user, entitiy.User{
			ID:    user1[i].ID,
			Name:  user1[i].UserName,
			Email: user1[i].Email,
		})
	}

	return user, nil
}

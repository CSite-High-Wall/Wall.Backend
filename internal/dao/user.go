package dao

import (
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return UserDao{
		db: db,
	}
}

func (dao UserDao) CreateUser(userName string, password string) error {
	result := dao.db.Create(&model.User{
		UserName: userName,
		Password: password,
	})

	return result.Error
}

func (dao UserDao) FindUserByUserName(userName string) (model.User, error) {
	var user model.User
	result := dao.db.First(&user, "user_name = ?", userName)

	return user, result.Error
}

func (dao UserDao) FindUserByUserId(userId uuid.UUID) (model.User, error) {
	var user model.User
	result := dao.db.First(&user, "user_id = ?", userId)

	return user, result.Error
}

func (dao UserDao) UpdateTokenOfUser(userID uuid.UUID, token_identifier uuid.UUID) error {
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("TokenIdentifier", token_identifier)
	return result.Error
}

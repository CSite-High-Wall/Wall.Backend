package service

import (
	"errors"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	userDao dao.UserDao
}

func NewUserService(dao dao.UserDao) UserService {
	return UserService{
		userDao: dao,
	}
}

func (service UserService) RegisterUser(requestBody model.RegisterRequestJsonObject) error {
	return service.userDao.CreateUser(requestBody.UserName, requestBody.Password)
}

func (service UserService) FindUserByUserName(userName string) (model.User, error) {
	return service.userDao.FindUserByUserName(userName)
}

func (service UserService) FindUserByUserId(userId uuid.UUID) (model.User, error) {
	return service.userDao.FindUserByUserId(userId)
}

func (service UserService) ContainsUserName(userName string) (bool, error) {
	user, error := service.userDao.FindUserByUserName(userName)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user.UserName == userName, error
}

func (service UserService) GetUserInfoByUserId(userId uuid.UUID) (interface{}, error) {
	user, error := service.userDao.FindUserByUserId(userId)

	if error != nil {
		return nil, error
	} else {
		return model.UserInfoResponseJsonObject{
			UserId:        userId,
			UserName:      user.UserName,
			AvatarUrl:     user.AvatarUrl,
			CreatedAt:     user.CreatedAt,
			LastLoginTime: user.LastLoginTime,
		}, nil
	}
}

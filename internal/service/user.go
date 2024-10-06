package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"
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
			NickName:      user.NickName,
			AvatarUrl:     user.AvatarUrl,
			CreatedAt:     user.CreatedAt,
			LastLoginTime: user.LastLoginTime,
		}, nil
	}
}

func (service UserService) UpdateNickName(userId uuid.UUID, nickName string) error {
	return service.userDao.UpdateNickNameOfUser(userId, nickName)
}

func (service UserService) UploadUserAvatarUrl(userId uuid.UUID, avatarUrl string) error {
	return service.userDao.UpdateAvatarUrlOfUser(userId, avatarUrl)
}

func (service UserService) UpdatePassword(userID uuid.UUID, oldPassword string, newPassword string) error {
	user, err := service.userDao.FindUserByUserId(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("未找到该用户")
	} else if err != nil {
		return errors.New("获取用户信息失败")
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil { // 验证旧密码
		return errors.New("旧密码不正确")
	} else if regex := regexp.MustCompile(`^[a-zA-Z0-9@#$%^&*]{8,30}$`); !regex.MatchString(newPassword) {
		return errors.New("密码必须是：8-30位字符，只允许数字、大小写字母、以及 @ # $ % ^ & * 字符")
	} else {
		return service.userDao.UpdatePasswordOfUser(userID, newPassword) // 更新新密码
	}
}

package dao

import (
	"time"
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// 创建用户
func (dao UserDao) CreateUser(userName string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	result := dao.db.Create(&model.User{
		UserName: userName,
		Password: string(hash),
	})

	return result.Error
}

// 根据用户名查找用户，看是否已经存在
func (dao UserDao) FindUserByUserName(userName string) (model.User, error) {
	var user model.User
	result := dao.db.First(&user, "user_name = ?", userName)

	return user, result.Error
}

// 根据用户id查找用户，看是否已经存在
func (dao UserDao) FindUserByUserId(userId uuid.UUID) (model.User, error) {
	var user model.User
	result := dao.db.First(&user, "user_id = ?", userId)

	return user, result.Error
}

// 更新数据库中指定用户的令牌标识符
func (dao UserDao) UpdateTokenOfUser(userID uuid.UUID, tokenIdentifier uuid.UUID) error {
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("TokenIdentifier", tokenIdentifier)
	return result.Error
}

// 更新数据库中指定用户的最后一次登录时间
func (dao UserDao) UpdateLastLoginTimeOfUser(userID uuid.UUID) error {
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("LastLoginTime", time.Now())
	return result.Error
}

// 更新数据库中指定用户的昵称
func (dao UserDao) UpdateNickNameOfUser(userID uuid.UUID, nickName string) error {
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("NickName", nickName)
	return result.Error
}

// 更新数据库中指定用户的头像 Url
func (dao UserDao) UpdateAvatarUrlOfUser(userID uuid.UUID, avatarUrl string) error {
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("AvatarUrl", avatarUrl)
	return result.Error
}

// 更新数据库中指定用户的密码
func (dao UserDao) UpdatePasswordOfUser(userID uuid.UUID, newPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	result := dao.db.Model(model.User{}).Where("user_id = ?", userID).Update("Password", string(hash))
	return result.Error
}

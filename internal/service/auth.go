package service

import (
	"time"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret = []byte("73EEC309-0BE7-4A58-BE29-84307F3F56CD")

type AuthService struct {
	userDao dao.UserDao
}

func NewAuthService(userDao dao.UserDao) AuthService {
	return AuthService{
		userDao: userDao,
	}
}

func (service AuthService) Authenticate(userId uuid.UUID) (interface{}, error) {
	user, error := service.userDao.FindUserByUserId(userId)

	if error != nil {
		return nil, error
	}

	uuid := uuid.New()
	accessToken, error := service.GenerateAccessToken(user.UserId, uuid)

	if error != nil {
		return nil, error
	}

	if error := service.userDao.UpdateTokenOfUser(user.UserId, uuid); error != nil {
		return nil, error
	} else if error := service.userDao.UpdateLastLoginTimeOfUser(user.UserId); error != nil {
		return nil, error
	}

	return model.AuthTokenResponseJsonObject{
		UserId:      user.UserId,
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   21600,
	}, nil
}

func (service AuthService) Signout(userId uuid.UUID) error {
	user, error := service.userDao.FindUserByUserId(userId)

	if error != nil {
		return error
	}

	if error := service.userDao.UpdateTokenOfUser(user.UserId, uuid.Nil); error != nil {
		return error
	}

	return nil
}

func (service AuthService) GenerateAccessToken(userId uuid.UUID, token_identifier uuid.UUID) (string, error) {
	claims := model.UserClaims{
		UserId:          userId,
		TokenIdentifier: token_identifier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func (service AuthService) VerifyAccessToken(accessToken string) (bool, uuid.UUID) {
	if token, error := jwt.ParseWithClaims(accessToken, &model.UserClaims{}, GetJwtKeyFunc()); error != nil {
		return false, uuid.Nil
	} else {
		userClaims, ok := token.Claims.(*model.UserClaims)
		user, error := service.userDao.FindUserByUserId(userClaims.UserId)

		if error != nil {
			return false, uuid.Nil
		} else if !ok || user.TokenIdentifier == uuid.Nil {
			return false, uuid.Nil
		} else if user.TokenIdentifier != userClaims.TokenIdentifier {
			return false, uuid.Nil
		} else {
			if token.Valid {
				return true, userClaims.UserId
			} else {
				return false, uuid.Nil
			}
		}
	}
}

func GetJwtKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}
}

package main

import (
	"wall-backend/internal/controller"
	"wall-backend/internal/dao"
	"wall-backend/internal/middleware"
	"wall-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	InitComponents()
	InitRoute(r)

	r.Run(":8000")
}

var ConfigService service.ConfigService
var DataBaseService service.DataBaseService
var UserService service.UserService
var AuthService service.AuthService
var ExpressionService service.ExpressionService

var UserDao dao.UserDao
var ExpressionDao dao.ExpressionDao

var RegisterController controller.RegisterController
var AuthController controller.AuthController
var ProfileController controller.ProfileController
var ExpressController controller.ExpressController
var CommunityController controller.CommunityController
var ReviewController controller.ReviewController

func InitComponents() {
	ConfigService = service.NewConfigService()
	ConfigService.Initialize()

	DataBaseService = service.NewDataBaseService(ConfigService)
	DataBaseService.Connect()
	DataBaseService.InitializeDataTable()

	UserDao = dao.NewUserDao(DataBaseService.DB)
	ExpressionDao = dao.NewExpressionDao(DataBaseService.DB)

	UserService = service.NewUserService(UserDao)
	AuthService = service.NewAuthService(UserDao)
	ExpressionService = service.NewExpressionService(ExpressionDao)

	RegisterController = controller.NewRegisterController(UserService)
	AuthController = controller.NewAuthController(AuthService, UserService)
	ProfileController = controller.NewProfileController(UserService)
	ExpressController = controller.NewExpressController(UserService, ExpressionService)
	CommunityController = controller.NewCommunityController(UserService, ExpressionService)

	middleware.AuthService = AuthService
}

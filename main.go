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
var ReviewService service.ReviewService

var UserDao dao.UserDao
var ExpressionDao dao.ExpressionDao
var ReviewDao dao.ReviewDao

var RegisterController controller.RegisterController
var AuthController controller.AuthController
var ProfileController controller.ProfileController
var ExpressController controller.ExpressController
var ReviewController controller.ReviewController
var CommunityController controller.CommunityController

func InitComponents() {
	ConfigService = service.NewConfigService()
	ConfigService.Initialize()

	DataBaseService = service.NewDataBaseService(ConfigService)
	DataBaseService.Connect()
	DataBaseService.InitializeDataTable()

	UserDao = dao.NewUserDao(DataBaseService.DB)
	ExpressionDao = dao.NewExpressionDao(DataBaseService.DB)
	ReviewDao = dao.NewReviewDao(DataBaseService.DB)

	UserService = service.NewUserService(UserDao)
	AuthService = service.NewAuthService(UserDao)
	ExpressionService = service.NewExpressionService(ExpressionDao)
	ReviewService = service.NewReviewService(ReviewDao)

	RegisterController = controller.NewRegisterController(UserService)
	AuthController = controller.NewAuthController(AuthService, UserService)
	ProfileController = controller.NewProfileController(UserService, ExpressionService)
	ExpressController = controller.NewExpressController(UserService, ExpressionService)
	ReviewController = controller.NewReviewController(UserService, ReviewService, ExpressionService)
	CommunityController = controller.NewCommunityController(UserService, ExpressionService, ReviewService)

	middleware.AuthService = AuthService
}

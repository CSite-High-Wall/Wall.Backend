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

	error := r.Run(":8000")

	if error != nil {
		panic(error)
	}
}

var ConfigService service.ConfigService
var DataBaseService service.DataBaseService
var UserService service.UserService
var AuthService service.AuthService
var ExpressionService service.ExpressionService
var ReviewService service.ReviewService
var BlacklistService service.BlacklistService

var UserDao dao.UserDao
var ExpressionDao dao.ExpressionDao
var ReviewDao dao.ReviewDao
var BlacklistDao dao.BlacklistDao

var RegisterController controller.RegisterController
var AuthController controller.AuthController
var ProfileController controller.ProfileController
var ExpressController controller.ExpressController
var ReviewController controller.ReviewController
var CommunityController controller.CommunityController
var BlacklistController controller.BlacklistController

func InitComponents() {
	ConfigService = service.NewConfigService()
	if error := ConfigService.Initialize(); error != nil {
		panic(error)
	}

	DataBaseService = service.NewDataBaseService(ConfigService)
	if error := DataBaseService.Connect(); error != nil {
		panic(error)
	}
	if error := DataBaseService.InitializeDataTable(); error != nil {
		panic(error)
	}

	UserDao = dao.NewUserDao(DataBaseService.DB)
	ExpressionDao = dao.NewExpressionDao(DataBaseService.DB)
	ReviewDao = dao.NewReviewDao(DataBaseService.DB)
	BlacklistDao = dao.NewBlacklistDao(DataBaseService.DB)

	UserService = service.NewUserService(UserDao)
	AuthService = service.NewAuthService(UserDao)
	ExpressionService = service.NewExpressionService(ExpressionDao)
	ReviewService = service.NewReviewService(ReviewDao)
	BlacklistService = service.NewBlacklistService(BlacklistDao)

	RegisterController = controller.NewRegisterController(UserService)
	AuthController = controller.NewAuthController(AuthService, UserService)
	ProfileController = controller.NewProfileController(UserService, ExpressionService, AuthService, ConfigService)
	ExpressController = controller.NewExpressController(UserService, ExpressionService)
	ReviewController = controller.NewReviewController(UserService, ReviewService, ExpressionService)
	CommunityController = controller.NewCommunityController(UserService, ExpressionService, ReviewService, BlacklistService, AuthService)
	BlacklistController = controller.NewBlacklistController(UserService, BlacklistService, ExpressionService)

	middleware.AuthService = AuthService
}

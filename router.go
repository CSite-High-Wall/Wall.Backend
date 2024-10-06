package main

import (
	"net/http"
	"wall-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	r.Use(middleware.CorsHandler)

	api := r.Group("/api")
	{
		api.POST("register", RegisterController.Register)

		authserver := api.Group("/authserver")
		authserver.POST("/authenticate", AuthController.Authenticate)
		authserver.POST("/refresh", AuthController.Refresh)
		authserver.POST("/signout", AuthController.Signout)
		authserver.POST("/validate", AuthController.Validate)

		profile := api.Group("/profile")
		profile.GET("/user-info", middleware.AuthToken, ProfileController.GetUserInfo)
		profile.GET("/expressions", middleware.AuthToken, ProfileController.FetchUserExpressions)
		profile.POST("/nickname/edit", middleware.AuthToken, ProfileController.EditNickName)
		profile.PUT("/avatar/upload", middleware.AuthToken, ProfileController.UploadUserAvatar)
		profile.POST("/password/change", middleware.AuthToken, ProfileController.ChangePassword)

		userBlacklist := profile.Group("/user-blacklist")
		userBlacklist.POST("/add", middleware.AuthToken, BlacklistController.AddUserIntoBlacklist)
		userBlacklist.DELETE("/remove", middleware.AuthToken, BlacklistController.RemoveUserFromBlacklist)
		userBlacklist.GET("/get", middleware.AuthToken, BlacklistController.GetUserBlacklist)

		expressionBlacklist := profile.Group("/expression-blacklist")
		expressionBlacklist.POST("/add", middleware.AuthToken, BlacklistController.AddExpressionIntoBlacklist)
		expressionBlacklist.DELETE("/remove", middleware.AuthToken, BlacklistController.RemoveExpressionFromBlacklist)
		expressionBlacklist.GET("/get", middleware.AuthToken, BlacklistController.GetExpressionBlacklist)

		community := api.Group("/community")
		community.GET("/expressions", CommunityController.FetchAllExpression)
		community.GET("/expression", CommunityController.FetchTargetedExpression)
		community.GET("/review", CommunityController.FetchAllReviewOfExpression)

		expression := api.Group("/express")
		expression.POST("/publish", middleware.AuthToken, ExpressController.Publish)
		expression.PUT("/edit", middleware.AuthToken, ExpressController.Edit)
		expression.DELETE("/delete", middleware.AuthToken, ExpressController.Delete)

		review := api.Group("/review")
		review.POST("/publish", middleware.AuthToken, ReviewController.Publish)
		review.DELETE("/delete", middleware.AuthToken, ReviewController.Delete)
		review.PUT("/edit", middleware.AuthToken, ReviewController.Edit)
		// review.POST("/reply", middleware.AuthToken, ReviewController.Reply)

		api.StaticFS("/static", http.Dir("static"))
	}

	r.NoMethod(middleware.NotFoundHandler)
	r.NoRoute(middleware.NotFoundHandler)
}

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
		// profile.POST("/username/edit", middleware.AuthToken, ProfileController.EditUserName)
		profile.PUT("/avatar/upload", middleware.AuthToken, ProfileController.UploadUserAvatar)
		//profile.POST("/avatar/upload", middleware.AuthToken, ProfileController.UploadUserAvatarUrl)
		profile.PUT("/password/edit",middleware.AuthToken,ProfileController.UpdatePassword)

		blacklist := profile.Group("/blacklist")
		blacklist.POST("/add", middleware.AuthToken, BlacklistController.Add)
		blacklist.DELETE("/remove", middleware.AuthToken, BlacklistController.Remove)
		blacklist.GET("/get", middleware.AuthToken, BlacklistController.GetBlacklistOfUser)

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

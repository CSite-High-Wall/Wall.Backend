package main

import (
	"wall-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	r.Use(middleware.CorsHandler)
	r.NoMethod(middleware.NotFoundHandler)
	r.NoRoute(middleware.NotFoundHandler)

	api := r.Group("/api")
	{
		api.POST("register", RegisterController.Register)

		authserver := api.Group("/authserver")
		authserver.POST("/authenticate", AuthController.Authenticate)
		authserver.POST("/refersh", AuthController.Refresh)
		authserver.POST("/signout", AuthController.Signout)
		authserver.POST("/validate", AuthController.Validate)

		// b := api.Group("/person")
		// {
		// 	b.POST("/nickname/edit", UserController.Nickname)
		// }

		expression := api.Group("/express")
		expression.PUT("/edit", middleware.AuthToken, ExpressController.Edit)
		expression.DELETE("/delete", middleware.AuthToken, ExpressController.Delete)
		expression.POST("/publish", middleware.AuthToken, ExpressController.Publish)

		community := api.Group("/community")
		community.GET("/expressions", CommunityController.FetchAllExpression)
		community.GET("/expression", CommunityController.FetchTargetedExpression)

		review := api.Group("/review")
		review.POST("/publish", ReviewController.Publish)
		review.DELETE("/delete", ReviewController.Delete)
		review.PUT("/edit", ReviewController.Edit)

		// person:= api.Group("/profile")
		// blacklist:= person.Group("/blacklist")
		// blacklist.POST("/add", middleware.AuthToken, BlacklistController.Add)
		// blacklist.DELETE("/remove", middleware.AuthToken, BlacklistController.Remove)
		// blacklist.GET("/get", middleware.AuthToken, BlacklistController.Get)

		// review.POST("/reply", ReviewController.Reply)

		// api.POST("login", user.Login)

		// c := api.Group("/contact")
		// {
		// 	c.POST("", contact.CreateContact)
		// 	c.PUT("", contact.UpdateContact)
		// 	c.DELETE("", contact.DeleteContact)
		// 	c.GET("", contact.GetContact)
		// }
	}
}

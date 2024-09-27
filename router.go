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

		// b := api.Group("/person")
		// {
		// 	b.POST("/nickname/edit", UserController.Nickname)
		// }
		d := api.Group("/expression")
		{
			d.PUT("/edit", ExpressController.Edit)
			d.DELETE("/delete", ExpressController.Delete)
			d.POST("/publish", ExpressController.Publish)
		}
		f := api.Group("/review")
		{
			f.POST("/publish", ReviewController.Publish)
			f.DELETE("/delete", ReviewController.Delete)
			f.PUT("/edit", ReviewController.Edit)
			// f.POST("/reply", ReviewController.Reply)
		}

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

package controller

import "github.com/gin-gonic/gin"

type ReviewController struct {
}

func NewReviewController() ReviewController {
	return ReviewController{}
}

func (controller ReviewController) Publish(c *gin.Context) {

}

func (controller ReviewController) Delete(c *gin.Context) {

}

func (controller ReviewController) Edit(c *gin.Context) {

}

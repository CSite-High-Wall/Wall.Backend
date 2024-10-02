package controller

import (
	"github.com/google/uuid"
	"strconv"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CommunityController struct {
	userService       service.UserService
	expressionService service.ExpressionService
	reviewService     service.ReviewService
	blacklistService  service.BlacklistService
}

func NewCommunityController(userService service.UserService, expressionService service.ExpressionService, reviewService service.ReviewService, blacklistService service.BlacklistService) CommunityController {
	return CommunityController{
		userService:       userService,
		expressionService: expressionService,
		reviewService:     reviewService,
		blacklistService:  blacklistService,
	}
}

// 获取所有表白
func (controller CommunityController) FetchAllExpression(c *gin.Context) {
	var userId uuid.UUID = uuid.Nil
	var enableTruncation bool = false
	if _userId, isUserIdExist := c.GetQuery("user_id"); isUserIdExist {
		userId, _ = uuid.Parse(_userId)
	}
	if _enableTruncation, isUserIdExist := c.GetQuery("enable_truncation"); isUserIdExist {
		enableTruncation, _ = strconv.ParseBool(_enableTruncation)
	}

	expressions, err := controller.expressionService.FetchAllExpression()

	if err != nil {
		utils.ResponseFailWithoutData(c, "获取表白列表失败") // 如果查询出错，返回内部服务器错误
	} else {
		var expressionList []gin.H               // 准备响应数据
		for _, expression := range expressions { // 遍历所有表白，将每个表白的信息添加到expressionList中

			user, error := controller.userService.FindUserByUserId(expression.UserId)

			if error != nil {
				continue
			}

			var displayUserName string = "匿名用户"
			var displayAvatar string = ""

			if !expression.Anonymity {
				displayUserName = user.UserName
				displayAvatar = user.AvatarUrl
			}

			if enableTruncation {
				expression.Content = utils.TruncateText(expression.Content, 200)
			}

			expressionList = append(expressionList, gin.H{
				"expression_id": expression.ExpressionId,
				"user_id":       expression.UserId,
				"user_name":     displayUserName,
				"avatar_url":    displayAvatar,
				"title":         expression.Title,
				"content":       expression.Content,
				"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}

		if len(expressionList) == 0 {
			utils.ResponseOk(c, gin.H{
				"expression_list": [0]gin.H{}, // 准备最终响应
			}) // 返回成功响应，包含所有表白信息
		} else {
			if userId != uuid.Nil {
				filteredList, error := controller.blacklistService.FilterUserInBlacklist(userId, expressionList)

				if error != nil {
					utils.ResponseFailWithoutData(c, "拉取表白评论失败")
				} else {
					expressionList = filteredList
				}
			}

			utils.ResponseOk(c, gin.H{
				"expression_list": expressionList, // 准备最终响应
			}) // 返回成功响应，包含所有表白信息
		}
	}
}

// 获取指定表白
func (controller CommunityController) FetchTargetedExpression(c *gin.Context) {
	var expressionId uint64 = 0

	if expression_id, isUserIdExist := c.GetQuery("expression_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		expressionId, _ = strconv.ParseUint(expression_id, 10, 32)
	}

	expression, err := controller.expressionService.FindExpressionByExpressionId(expressionId)
	if err != nil {
		utils.ResponseFailWithoutData(c, "获取指定表白失败") // 如果查询出错，返回内部服务器错误
	} else {
		user, error := controller.userService.FindUserByUserId(expression.UserId)
		if error != nil {
			utils.ResponseFailWithoutData(c, "获取指定表白失败")
		} else {
			var displayUserName string = "匿名用户"
			var displayAvatar string = ""

			if !expression.Anonymity {
				displayUserName = user.UserName
				displayAvatar = user.AvatarUrl
			}

			utils.ResponseOk(c, gin.H{
				"expression_id": expression.ExpressionId,
				"user_id":       expression.UserId,
				"user_name":     displayUserName,
				"avatar_url":    displayAvatar,
				"content":       expression.Content,
				"title":         expression.Title,
				"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}
	}
}

// 获取表白下的所有评论
func (controller CommunityController) FetchAllReviewOfExpression(c *gin.Context) {
	var expressionId uint64 = 0
	var userId uuid.UUID = uuid.Nil

	if _userId, isUserIdExist := c.GetQuery("user_id"); isUserIdExist {
		userId, _ = uuid.Parse(_userId)
	}
	if expression_id, isUserIdExist := c.GetQuery("expression_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		expressionId, _ = strconv.ParseUint(expression_id, 10, 32)
	}

	reviews, error := controller.reviewService.FindReviewByExpressionId(expressionId)

	if error != nil {
		utils.ResponseFailWithoutData(c, "拉取表白评论失败") // 如果查询出错，返回内部服务器错误
	} else {
		var reviewList []gin.H
		for _, review := range reviews {
			user, error := controller.userService.FindUserByUserId(review.UserId)
			if error != nil {
				continue
			}

			reviewList = append(reviewList, gin.H{
				"expression_id": review.ExpressionId,
				"review_id":     review.ReviewId,
				"user_id":       review.UserId,
				"user_name":     user.UserName,
				"avatar_url":    user.AvatarUrl,
				"content":       review.Content,
				"created_at":    review.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}

		if len(reviewList) == 0 {
			utils.ResponseOk(c, gin.H{
				"review_list": [0]gin.H{}, // 准备最终响应
			})
		} else {
			if userId != uuid.Nil {
				filteredList, error := controller.blacklistService.FilterUserInBlacklist(userId, reviewList)

				if error != nil {
					utils.ResponseFailWithoutData(c, "拉取表白评论失败")
				} else {
					reviewList = filteredList
				}
			}

			utils.ResponseOk(c, gin.H{
				"review_list": reviewList, // 准备最终响应
			})
		}
	}
}

package router

import (
	"byte_douyin/controller/comment"
	"byte_douyin/controller/user_info"
	"byte_douyin/controller/user_login"
	"byte_douyin/controller/video"
	"byte_douyin/middleware"
	"byte_douyin/models"

	"github.com/gin-gonic/gin"
)

func Init_Router() *gin.Engine {
	models.Init_DB()
	r := gin.Default()

	r.Static("static", "./static")

	apiRouter := r.Group("/douyin")

	//basic apis
	apiRouter.GET("/feed/", video.FeedVideoListHandler)
	apiRouter.GET("/user/", middleware.JWTMiddleWare(), user_info.UserInfoHandler)
	apiRouter.POST("/user/login/", middleware.SHAMiddleWare(), user_login.UserLoginHandler)
	apiRouter.POST("/user/register/", middleware.SHAMiddleWare(), user_login.UserRegisterHandler)
	apiRouter.POST("/publish/action/", middleware.JWTMiddleWare(), video.PublishVideoHandler)
	apiRouter.GET("/publish/list/", middleware.NoAuthToGetUserId(), video.QueryVideoListHandler)

	//extend api 1
	apiRouter.POST("/favorite/action/", middleware.JWTMiddleWare(), video.PostFavorHandler)
	apiRouter.GET("/favorite/list/", middleware.NoAuthToGetUserId(), video.QueryFavorVideoListHandler)
	apiRouter.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PostCommentHandler)
	apiRouter.GET("/comment/list/", middleware.JWTMiddleWare(), comment.QueryCommentListHandler)

	//extend api 2
	apiRouter.POST("/relation/action/", middleware.JWTMiddleWare(), user_info.PostFollowActionHandler)
	apiRouter.GET("/relation/follow/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowListHandler)
	apiRouter.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowerHandler)
	return r
}

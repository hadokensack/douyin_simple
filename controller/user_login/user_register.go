package user_login

import (
	"byte_douyin/models"
	"byte_douyin/service/user_login"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisterResponse struct {
	models.CommonResponse
	*user_login.LoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析出错",
			},
		})
		return
	}
	registerResponse, err := user_login.PostUserLogin(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginResponse:  registerResponse,
	})
}

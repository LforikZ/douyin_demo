package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]entity.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

var userIdSequence = int64(1)

func Register(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, entity.UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorParams},
		})
	}
}

func Login(c *gin.Context) {

	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, entity.UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorParams},
		})
	}

}

func UserInfo(c *gin.Context) {
	token := c.Query("token") //get请求，url截取token参数
	var claim *util.Claims
	if token == "" {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "错误: 没找到token!(" + e.ErrorNotToken + ")"})
		return
	} else {
		claims, err := util.ParseToken(token)
		claim = claims
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "解析token失败QAQ(" + e.ErrorAuthCheckTokenFail + " or " + e.ErrorAuth + ")"})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "WARN：token未在有效期QAQ(" + e.ErrorAuthCheckTokenTimeout + ")"})
			return
		}
	} //token分析

	var infoService service.InfoService
	res := infoService.Info(claim.Id, claim.Username)
	if res.StatusCode == 1 {
		c.JSON(http.StatusBadRequest, res.Response)
	} else {
		c.JSON(http.StatusOK, res.User)
	}

}

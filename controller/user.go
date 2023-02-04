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

//type UserRegisterResponse struct {
//	Response
//	UserId int64  `json:"user_id,omitempty"`
//	Token  string `json:"token"`
//}
//
//type UserLoginResponse struct {
//	Response
//	UserId int64  `json:"user_id,omitempty"`
//	Token  string `json:"token"`
//}
//
//type UserResponse struct {
//	Response
//	User entity.User `json:"user"`
//}

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
	//username := c.Query("username")
	//password := c.Query("password")
	//
	//token := username + password
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//} else {
	//	atomic.AddInt64(&userIdSequence, 1)
	//	newUser := entity.User{
	//		Id:   userIdSequence,
	//		Name: username,
	//	}
	//	usersLoginInfo[token] = newUser
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   userIdSequence,
	//		Token:    username + password,
	//	})
	//}
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
	//username := c.Query("username")
	//password := c.Query("password")
	//
	//token := username + password
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   user.Id,
	//		Token:    token,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 41, StatusMsg: "错误: 没找到token!(" + e.ErrorNotToken + ")"})
		return
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 41, StatusMsg: "解析token失败QAQ(" + e.ErrorAuthCheckTokenFail + " or " + e.ErrorAuth + ")"})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 41, StatusMsg: "WARN：token未在有效期QAQ(" + e.ErrorAuthCheckTokenTimeout + ")"})
			return
		}
	} //token分析

	// TODO 查数据库，正在写 2023-02-04 22：01：20
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 51, StatusMsg: e.ErrorUserNotFound},
		})
	}
	//token := c.Query("token")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK,{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	var infoService service.InfoService
	id := c.Query("id") //get请求，url截取id参数
	if id != "" {
		userId, err := strconv.Atoi(id)
		if err == nil {
			res := infoService.InfoByID(uint(userId))
			if res.StatusCode == 1 {
				c.JSON(http.StatusBadRequest, res.Response)
			} else {
				c.JSON(http.StatusOK, res.User)
			}
		}
	}

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

	res := infoService.InfoByToken(claim.Id, claim.Username)
	if res.StatusCode == 1 {
		c.JSON(http.StatusBadRequest, res.Response)
	} else {
		c.JSON(http.StatusOK, res.User)
	}

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: entity.Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorUserNotFound},
	//	})
	//}
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

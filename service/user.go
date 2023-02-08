// @Author Junyi Tan 2023/2/3 23:30:00
package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"gorm.io/gorm"
)

type UserService struct {
	Name     string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

type InfoService struct {
	Id   int64  `form:"id" json:"id" binding:"required"`
	Name string `form:"username" json:"username" binding:"required,max=32"`
}

func (service *UserService) Register() *entity.UserRegisterResponse {
	var code = e.CodeFailed
	var user mysql.User
	var count int64
	count = mysql.RegisterAuth(service.Name, &user)

	// 表单验证
	if count == 1 {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorExistUser},
			Token:    "",
		}
	}
	user.Uid = util.GetID()
	user.Name = service.Name
	user.FollowerCount = 0
	user.FollowCount = 0
	user.IsFollow = false

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: code, StatusMsg: e.ErrorFailEncryption},
			Token:    "",
		}
	}
	// 创建用户
	err := mysql.CreateUser(&user)
	if err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: code, StatusMsg: e.ErrorDatabase},
			Token:    "",
		}
	}
	//生成token
	token, err := util.GenerateToken(user.Uid, user.Name, user.FollowCount,
		user.FollowerCount, user.IsFollow, 0)
	if err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{
				StatusCode: code,
				StatusMsg:  e.ErrorAuthToken,
			},
			Token: "",
		}
	}

	code = e.CodeSuccess
	return &entity.UserRegisterResponse{
		Response: entity.Response{StatusCode: code, StatusMsg: e.RegisterSuccess},
		UserId:   user.Uid,
		Token:    token,
	}
}

func (service *UserService) Login() *entity.UserRegisterResponse {
	var user mysql.User
	code := e.CodeFailed
	count := mysql.RegisterAuth(service.Name, &user)

	// 表单验证
	if count == 0 {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorNotExistUser},
			Token:    "",
		}
	}
	if err := mysql.LoginAuth(service.Name, &user); err != nil {
		// 如果查询不到，返回相应的错误
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return &entity.UserRegisterResponse{
				Response: entity.Response{
					StatusCode: code,
					StatusMsg:  e.ErrorUserNotFound,
				},
				Token: "",
			}
		}
		return &entity.UserRegisterResponse{
			Response: entity.Response{
				StatusCode: code,
				StatusMsg:  e.ErrorDatabase,
			},
			Token: "",
		}

	}
	if !user.CheckPassword(service.Password) {
		return &entity.UserRegisterResponse{
			Response: entity.Response{
				StatusCode: code,
				StatusMsg:  e.ErrorNotCompare,
			},
			Token: "",
		}
	}
	token, err := util.GenerateToken(user.Uid, user.Name, user.FollowCount,
		user.FollowerCount, user.IsFollow, 0)
	if err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{
				StatusCode: code,
				StatusMsg:  e.ErrorAuthToken,
			},
			Token: "",
		}
	}
	code = e.CodeSuccess
	return &entity.UserRegisterResponse{
		Response: entity.Response{
			StatusCode: code,
			StatusMsg:  e.LoginSuccess,
		},
		UserId: user.Uid,
		Token:  token,
	}
}

func (service *InfoService) Info(uid int64, name string) *entity.UserResponse {
	var user mysql.User
	err := mysql.InfoAuth(&user, uid)
	if err != nil {
		user := entity.User{Name: name, Id: service.Id}
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return &entity.UserResponse{
				Response: entity.Response{
					StatusCode: 1,
					StatusMsg:  e.ErrorUserNotFound,
				},
				User: user,
			}
		}
		return &entity.UserResponse{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorDatabase,
			},
			User: user,
		}
	}

	code := e.CodeSuccess
	return &entity.UserResponse{
		Response: entity.Response{
			StatusCode: code,
			StatusMsg:  e.UserSelectSuccess,
		},
		User: entity.User{Name: user.Name, Id: user.Uid,
			FollowCount: user.FollowCount, FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow},
	}
}

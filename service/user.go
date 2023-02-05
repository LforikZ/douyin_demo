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
	id   int64  `form:"id" json:"id" binding:"required"`
	Name string `form:"username" json:"username" binding:"required,max=32"`
}

func (service *UserService) Register() *entity.UserRegisterResponse {
	var code = e.CodeFailed
	var user mysql.User
	var count int64
	count = mysql.RegisterAuth((*mysql.UserService)(service), &user)

	// 表单验证
	if count == 1 {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: e.ErrorExistUser},
			Token:    "",
		}
	}
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
	//生成token
	token, err := util.GenerateToken(user.ID, service.Name, 0)
	if err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: code, StatusMsg: e.ErrorAuthToken},
			Token:    "",
		}
	}
	// 创建用户
	err = mysql.CreateUser(&user)
	if err != nil {
		return &entity.UserRegisterResponse{
			Response: entity.Response{StatusCode: code, StatusMsg: e.ErrorDatabase},
			Token:    "",
		}
	}
	code = e.CodeSuccess
	return &entity.UserRegisterResponse{
		Response: entity.Response{StatusCode: code, StatusMsg: e.RegisterSuccess},
		UserId:   int64(user.ID),
		Token:    token,
	}
}

func (service *UserService) Login() *entity.UserRegisterResponse {
	var user mysql.User
	code := e.CodeFailed
	if err := mysql.LoginAuth((*mysql.UserService)(service), &user); err != nil {
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
	token, err := util.GenerateToken(user.ID, service.Name, 0)
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
		UserId: int64(user.ID),
		Token:  token,
	}
}

func (service *InfoService) Info(id uint, name string) *entity.UserResponse {
	var user mysql.User
	err := mysql.InfoAuth(&user, id)
	if err != nil {
		user := entity.User{Name: name, Id: service.id}
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
		User: entity.User{Name: user.Name, Id: int64(user.ID),
			FollowCount: user.FollowCount, FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow},
	}
}

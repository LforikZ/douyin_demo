package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// CommentAction 登录用户对视频进行评论
func CommentAction(c *gin.Context) {
	// 获取参数数据并进行校验
	p := new(entity.ParamComment)
	if err := c.ShouldBindQuery(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeFailed,
				StatusMsg:  ParamsError,
			})
			return
		}
		errData := RemoveTopStruct(errs.Translate(Trans)) // 翻译并去除错误中结构体名字
		c.JSON(http.StatusOK, ResponseValim{
			Response: Response{
				StatusCode: CodeFailed,
				StatusMsg:  ValidatorError,
			},
			Data: errData,
		})
		return
	}

	if p.ActionType == 1 && p.Content == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  ContentLoss,
		})
		return
	}
	if p.ActionType == 2 && p.CommentID == 0 {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  CommentIDLoss,
		})
		return
	}
	//验证token
	authentication, s := util.Authentication(p.Token)
	if authentication == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  s,
		})
		return
	}
	// 验证视频id
	result, err := service.GetVideoByVid(p.VideoID)
	if err != nil || result == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  VideosNotExit,
		})
		return
	}
	// 业务处理
	comment, err := service.InsertComment(p)
	if err != nil {
		if p.ActionType == 1 {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeFailed,
				StatusMsg:  InsertCommentError,
			})
			return
		} else if p.ActionType == 2 {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeFailed,
				StatusMsg:  DeleteCommentError,
			})
			return
		}
	} else {
		if p.ActionType == 1 {
			c.JSON(http.StatusOK, CommentResponse{
				Response: Response{
					StatusCode: CodeSuccess,
					StatusMsg:  InsertCommentSuccess,
				},
				Comment: *comment,
			})
			return
		} else if p.ActionType == 2 {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeSuccess,
				StatusMsg:  DeleteCommentSuccess,
			})
			return
		}
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	// 获取参数数据并进行校验
	p := new(entity.ParamTokenVID)
	if err := c.ShouldBindQuery(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeFailed,
				StatusMsg:  ParamsError,
			})
			return
		}
		errData := RemoveTopStruct(errs.Translate(Trans)) // 翻译并去除错误中结构体名字
		c.JSON(http.StatusOK, ResponseValim{
			Response: Response{
				StatusCode: CodeFailed,
				StatusMsg:  ValidatorError,
			},
			Data: errData,
		})
		return
	}

	// 业务处理
	list, err := service.GetCommentList(p)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  GetCommentListFiled,
		})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: CodeSuccess},
		CommentList: list,
	})
	return
}

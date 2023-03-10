package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"path/filepath"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 数据校验
	token := c.PostForm("token")
	userInfo, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  UserNotExit,
		})
	}
	user := entity.User{
		Id:            userInfo.Id,
		Name:          userInfo.Username,
		FollowCount:   userInfo.FollowerCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 业务处理
	saveFile, finalName := service.Publish(data, user)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  err.Error(),
		})
		return
	}
	info, err := GetVideoInfo(finalName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  err.Error(),
		})
	}

	if err := service.InsertVideo(info, user); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: CodeSuccess,
		StatusMsg:  finalName + UploadSuccess,
	})
	return
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// 获取参数，并判断参数是否有效
	var p entity.ParamTokenUID
	if err := c.ShouldBindQuery(&p); err != nil {
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
	// 验证token是否有效
	authentication, s := util.Authentication(p.Token)
	if authentication == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  s,
		})
		return
	}

	videoList, err := service.GetVideoList(p.UserID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  VideosNotExit,
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: CodeSuccess,
		},
		VideoList: videoList,
	})
	return
}

// GetVideoInfo 解析视频内容
func GetVideoInfo(videoPath string) (os.FileInfo, error) {
	return os.Stat(filepath.Join("./public/", videoPath))
}

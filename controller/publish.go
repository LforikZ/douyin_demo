package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []entity.ApiVideo `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 数据校验
	token := c.PostForm("token")
	//TODO: 等登录注册功能实现之后  修改验证 token 操作
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: CodeFailed, StatusMsg: UserNotExit})
		return
	}
	user := usersLoginInfo[token]
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
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: CodeSuccess,
		},
		VideoList: nil,
	})
}

// GetVideoInfo 解析视频内容
func GetVideoInfo(videoPath string) (os.FileInfo, error) {
	return os.Stat(filepath.Join("./public/", videoPath))
}

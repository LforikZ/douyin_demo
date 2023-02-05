// @Author Zihao_Li 2023/2/3 14:11:00
package service

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

// Publish
// @Description 上传视频
// @Author Zihao_Li 2023-02-05 13:28:11
func Publish(data *multipart.FileHeader, user entity.User) (string, string) {
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	return saveFile, finalName
}

// InsertVideo
// @Description  将视频信息保存到数据库
// @Author Zihao_Li 2023-02-05 13:28:42
func InsertVideo(info os.FileInfo, user entity.User) (err error) {
	path := info.Name()

	video := &entity.Video{
		AuthorID:      strconv.FormatInt(user.Id, 10),
		AuthorName:    user.Name,
		PlayUrl:       path,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}

	err = mysql.Insert(video)
	if err != nil {
		return err
	}
	return
}

// GetVideoList
// @Description 获取视频列表
// @Author Zihao_Li 2023-02-05 13:28:57
func GetVideoList(userid uint) (videos []entity.ApiVideo, err error) {
	result, _ := mysql.GetUserAllVideos(userid)
	if result == nil {
		err = errors.New("videos not exit")
		return videos, err
	}

	for _, video := range result {
		//TODO: 调用方法：根据用户id获取用户信息
		// method()
		midData := entity.ApiVideo{
			User:          nil,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		}
		videos = append(videos, midData)
	}

	return videos, err
}

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
		AuthorID:      int(user.Id),
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
func GetVideoList(userstr string) (videos []entity.ApiVideo, err error) {
	userid, err := strconv.Atoi(userstr)
	if err != nil {
		return nil, err
	}
	user, err := mysql.GetUserInfo(userid)
	if err != nil {
		err = errors.New("user not exit")
		return nil, err
	}
	result, _ := mysql.GetUserAllVideos(user.Name)
	if result == nil {
		err = errors.New("videos not exit")
		return videos, err
	}

	for _, video := range result {
		midData := entity.ApiVideo{
			User:          user,
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

// GetVideoByVid
// @Description  根据视频id获取视频信息
// @Author Zihao_Li 2023-02-07 13:46:45
func GetVideoByVid(videoid int64) (video *mysql.Video, err error) {
	video, err = mysql.GetVideoByVid(videoid)
	if err != nil {
		return nil, err
	}
	return video, err
}

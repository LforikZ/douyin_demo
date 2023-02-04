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

func Publish(data *multipart.FileHeader, user entity.User) (string, string) {
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	return saveFile, finalName
}

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

func GetVideoList(userid string) (videos []entity.ApiVideo, err error) {
	result, _ := mysql.GetUserAllVideos(userid)
	//TODO: 调用方法：根据用户id获取用户信息
	// method()
	if result == nil {
		err = errors.New("videos not exit")
		return videos, err
	}
	return videos, err
}

package service

import (
	"github.com/RaymondCode/simple-demo/entity"

	"github.com/RaymondCode/simple-demo/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Feed(userId int64, latestTime int64) []entity.VideoInfo {
	var videos []mysql.Video
	Db.Model(&entity.Video{}).Where("create_time<?", latestTime).Order("create_time DESC").Limit(30).Find(&videos)

	var videoLists []entity.VideoInfo
	for _, video := range videos {
		var author entity.User
		Db.Model(&entity.User{}).Where("id=?", video.AuthorId).First(&author)

		videoLists = append(videoLists, entity.VideoInfo{
			Id:            video.ID,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	return videoLists
}

package service

import (
	"github.com/RaymondCode/simple-demo/entity"

	"github.com/RaymondCode/simple-demo/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Feed(userId int64, latestTime int64) []entity.ApiVideo {
	var videos []mysql.Video
	Db.Model(&entity.Video{}).Where("create_time<?", latestTime).Order("create_time DESC").Limit(30).Find(&videos)

	var videoLists []entity.ApiVideo
	for _, video := range videos {
		var author entity.User
		Db.Model(&entity.User{}).Where("id=?", video.AuthorID).First(&author)

		videoLists = append(videoLists, entity.ApiVideo{
			Id:            int64(video.ID),
			User:          &author,
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

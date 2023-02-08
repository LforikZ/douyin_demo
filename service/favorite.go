package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
)

func FavoriteList(userId string) []entity.ApiVideo {
	var videoIds []int64
	// 根据userId在favorite表里查找用户喜欢的videoIds
	Db.Model(&mysql.Favorite{}).Select("video_id").Where("user_id=?", userId).Find(videoIds)
	var videoLists []entity.ApiVideo
	for _, videoId := range videoIds {
		var video entity.Video
		var author entity.User
		// 根据videoId在video表里查找video视频信息
		Db.Model(&mysql.Video{}).Where("id=?", videoId).First(&video)
		// 根据视频的AuthorId在user表里查找author信息
		Db.Model(&mysql.User{}).Where("uid=?", video.AuthorID).First(&author)
		// 在follows表查询用户是否关注author
		var isFollow int64
		Db.Model(&mysql.Follows{}).Where("uid=? and to_user_id", userId, author.Id).Count(&isFollow)
		if isFollow != 0 {
			author.IsFollow = true
		} else {
			author.IsFollow = false
		}

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

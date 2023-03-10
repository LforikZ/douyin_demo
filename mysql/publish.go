// @Author Zihao_Li 2023/2/3 14:28:00
package mysql

import (
	"database/sql"

	"github.com/RaymondCode/simple-demo/pkg/util"

	"strconv"

	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID       int64  `gorm:"notnull"`       //视频id
	AuthorID      string `gorm:"notnull"`       //作者id
	Title         string `gorm:"column:title"`  //视频标题
	AuthorName    string `gorm:"notnull"`       //作者姓名
	PlayUrl       string `gorm:"notnull"`       //视频地址
	CoverUrl      string `gorm:"notnull"`       //封面网址
	FavoriteCount int64  `gorm:"default 0"`     //收藏计数
	CommentCount  int64  `gorm:"default 0"`     //评论计数
	IsFavorite    bool   `gorm:"default false"` //是否收藏
}

func Insert(video *entity.Video) (err error) {
	vi := Video{
		VideoID:       util.GenID(),
		AuthorID:      strconv.Itoa(video.AuthorID),
		AuthorName:    video.AuthorName,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	if result := db.Create(&vi); result.Error != nil {
		return result.Error
	}
	return
}

func GetUserAllVideos(userName string) (a []entity.ApiVideo, err error) {
	var videos []Video
	if result := db.Where("author_name=?", userName).Find(&videos); result.Error == sql.ErrNoRows {
		err = result.Error
	}
	for _, video := range videos {
		midData := entity.ApiVideo{
			User:          nil,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		}
		a = append(a, midData)
	}
	return a, err
}

func GetVideoByVid(videoID int64) (*Video, error) {
	var video *Video
	result := db.Where("video_id=?", videoID).Find(&video)
	return video, result.Error
}

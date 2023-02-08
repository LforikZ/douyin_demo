// @Author Zihao_Li 2023/2/7 12:58:00
package mysql

import (
	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CommentID   int64  // 评论id
	VideoID     int64  // 视频id
	UserID      int64  // 评论人ID
	CommentText string // 评论内容
}

func InsertComment(a *entity.Comment) (err error) {
	comment := Comment{
		CommentID:   a.CommentID,
		VideoID:     a.VideoID,
		UserID:      a.User.Id,
		CommentText: a.Content,
	}
	if result := db.Create(&comment); result.Error != nil {
		err = result.Error
	}
	return err
}

func DeleteComment(commentID int64) (err error) {
	var comment Comment
	if result := db.Where("comment_id=?", commentID).Delete(&comment); result.Error != nil {
		err = result.Error
	}
	return err
}

func GetVideoListByVid(videoID string) (a []entity.Comment, err error) {
	var comments []Comment
	if result := db.Where("video_id=?", videoID).Find(&comments); result.Error != nil {
		err = result.Error
	}
	for _, comment := range comments {
		midData := entity.Comment{
			CommentID: comment.CommentID,
			VideoID:   comment.VideoID,
			User: entity.User{
				Id:            comment.UserID,
				Name:          "",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
			Content:    comment.CommentText,
			CreateDate: comment.CreatedAt.Format("01-02"),
		}
		a = append(a, midData)
	}
	return a, err
}

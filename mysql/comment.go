// @Author Zihao_Li 2023/2/7 12:58:00
package mysql

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentID   int    // 评论id
	ActionType  int    // 1-发布评论，2-删除评论
	VideoID     int64  // 视频id
	UserID      int64  // 评论人ID
	CommentText string // 评论内容
}

func InsertComment() {

}

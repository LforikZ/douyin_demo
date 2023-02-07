// @Author Zihao_Li 2023/2/4 17:24:00
package entity

type ParamTokenUID struct {
	Token  string `form:"token" binding:"required"`
	UserID string `form:"user_id" binding:"required"`
}

type ParamComment struct {
	Token      string `form:"token"`                                    // 用户鉴权token
	VideoID    int64  `form:"video_id"`                                 // 视频id
	ActionType int    `form:"action_type" binding:"required,oneof=1 2"` // 1-发布评论，2-删除评论
	Content    string `form:"comment_text"`                             // 评论内容
	CommentID  int64  `form:"CommentID"`                                // 要删除的评论id，在action_type=2的时候使用
}

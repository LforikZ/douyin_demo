// @Author Zihao_Li 2023/2/3 14:53:00
package entity

type Comment struct {
	CommentID  int64  `json:"id"`
	VideoID    int64  `json:"video_id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

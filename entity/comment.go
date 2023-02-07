// @Author Zihao_Li 2023/2/3 14:53:00
package entity

type Comment struct {
	CommentID  int64 `json:"id,omitempty"`
	VideoID    int64
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

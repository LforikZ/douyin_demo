// @Author Zihao_Li 2023/2/3 14:52:00
package entity

// 返回给用户的视频信息
type VideoInfo struct {
	Id            int64  `json:"id,omitempty"` // 视频id，唯一标识
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

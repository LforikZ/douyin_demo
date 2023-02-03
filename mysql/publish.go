// @Author Zihao_Li 2023/2/3 14:28:00
package mysql

import "gorm.io/gorm"

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type Video struct {
	gorm.Model
	AuthorName    string `gorm:"notnull"` //作者姓名
	PlayUrl       string `gorm:"notnull"` //视频地址
	CoverUrl      string `gorm:"notnull"` //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
}

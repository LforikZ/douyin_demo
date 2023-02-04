// @Author Zihao_Li 2023/2/3 14:28:00
package mysql

import "gorm.io/gorm"

// 数据库映射的User模型
type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"follower_count"`
	IsFollow      bool   `gorm:"is_follow"`
}

// 数据库映射的Video模型
type Video struct {
	gorm.Model
	AuthorId      int64  `gorm:"notnull"` // 作者id
	AuthorName    string `gorm:"notnull"` //作者姓名
	PlayUrl       string `gorm:"notnull"` //视频地址
	CoverUrl      string `gorm:"notnull"` //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
	Title         string // 视频标题
}

package mysql

import "gorm.io/gorm"

// Favorite表，用户喜欢的视频
type Favorite struct {
	Id      int64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

type Follows struct {
	gorm.Model
	FollowId int64
	ToUserId int64
	IsFollow bool
}

// create table favorite
// (
// 	id bigint auto_increment primary key,
// 	user_id bigint NOT NULL,
// 	video_id bigint
// );

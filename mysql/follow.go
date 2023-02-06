package mysql

import (
	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	FollowId int64 //关注者Id
	ToUserId int64 //被关注者Id
	IsFollow bool
}

func FollowAuth(FollowId int64, ToUserId int64, follow *Follow) int64 {
	var count int64
	db.Model(Follow{}).Where("follow_id=? AND to_user_id=?",
		FollowId, ToUserId).Find(&follow).Count(&count)
	return count
}

func FollowCreate(follow *entity.Follow) error {
	var followdata = Follow{
		FollowId: follow.FollowId,
		ToUserId: follow.ToUserId,
		IsFollow: follow.IsFollow,
	}
	if err := db.Create(&followdata).Error; err != nil {
		return err
	}
	return nil
}

func FollowUpdate(follow entity.Follow) error {
	err := db.Model(&Follow{}).Where("follow_id=? AND to_user_id=?", follow.FollowId,
		follow.ToUserId).Update("is_follow", follow.IsFollow).Error
	return err
}

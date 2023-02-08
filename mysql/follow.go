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

//判断是否有这个用户
func IdAuth(id int64) bool {
	var user *User
	var count int64
	db.Model(User{}).Where("uid=?", id).Find(&user).Count(&count)
	return count == 1
}

//获取数据库中是否有两人的关注记录
func FollowAuth(FollowId int64, ToUserId int64, follow *Follow) (int64, bool) {
	var count int64
	db.Model(Follow{}).Where("follow_id=? AND to_user_id=?",
		FollowId, ToUserId).Find(&follow).Count(&count)
	return count, follow.IsFollow
}

//创建关注关系
func FollowCreate(follow *entity.Follow) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var followdata = Follow{
			FollowId: follow.FollowId,
			ToUserId: follow.ToUserId,
			IsFollow: follow.IsFollow,
		}
		if err := tx.Create(&followdata).Error; err != nil {
			return err
		}
		if err := FollowUserUpdate(*follow, true, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

//更新关注关系
func FollowUpdate(follow entity.Follow, oldfollow bool) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Follow{}).Where("follow_id=? AND to_user_id=?", follow.FollowId,
			follow.ToUserId).Update("is_follow", follow.IsFollow).Error; err != nil {
			return err
		}
		isFollow := follow.IsFollow
		if isFollow == true && oldfollow == false {
			FollowUserUpdate(follow, true, tx)
		}
		if isFollow == false && oldfollow == true {
			FollowUserUpdate(follow, true, tx)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// follow后更新user flag为true表示关注，否则为取消关注
func FollowUserUpdate(follow entity.Follow, flag bool, tx *gorm.DB) error {
	if flag == true {
		follow_id := follow.FollowId
		to_user_id := follow.ToUserId
		err := tx.Model(&User{}).Where("uid=?", follow_id).
			Update("follow_count", gorm.Expr("follow_count+ ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&User{}).Where("uid=?", to_user_id).
			Update("follower_count", gorm.Expr("follower_count+ ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		follow_id := follow.FollowId
		to_user_id := follow.ToUserId
		err := tx.Model(&User{}).Where("uid=?", follow_id).
			Update("follow_count", gorm.Expr("follow_count- ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&User{}).Where("uid=?", to_user_id).
			Update("follower_count", gorm.Expr("follower_count- ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	}

}

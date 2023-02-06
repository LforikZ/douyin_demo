package entity

type Follow struct {
	FollowId int64 //关注者Id
	ToUserId int64 //被关注者Id
	IsFollow bool  //是否关注
}

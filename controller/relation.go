package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/mysql"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	auth, msg := util.Authentication(token)
	if auth == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: e.CodeFailed,
			StatusMsg:  msg,
		})
		return
	}
	follow_id := util.GetUserId(token)
	toUserIdTmp := c.Query("to_user_id")
	to_user_id, err := strconv.ParseInt(toUserIdTmp, 10, 64)
	if err != nil {
		panic(err)
	}
	if mysql.IdAuth(to_user_id) == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: e.CodeFailed,
			StatusMsg:  e.ErrorUserNotFound,
		})
		return
	}
	var isFollow bool = false
	actionType := c.Query("action_type")
	if actionType == "1" {
		isFollow = true
	}
	var follow *mysql.Follow
	real_follow := entity.Follow{follow_id, to_user_id, isFollow}
	count, oldfollow := mysql.FollowAuth(follow_id, to_user_id, follow)
	fmt.Println(oldfollow)
	if count == 0 {
		err := mysql.FollowCreate(&entity.Follow{FollowId: follow_id, ToUserId: to_user_id, IsFollow: isFollow})
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: e.CodeFailed,
				StatusMsg:  msg,
			})
			return
		}
	} else {
		mysql.FollowUpdate(real_follow, oldfollow)

	}
	var statusfollow string
	if isFollow == true {
		statusfollow = e.FollowSuccess
	} else {
		statusfollow = e.UnfollowSuccess
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: e.CodeSuccess,
		StatusMsg:  statusfollow,
	})
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

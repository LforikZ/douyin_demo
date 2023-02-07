package controller

import (
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 解析传入的latest_time参数
	//var latestTime int64
	//if latestTimeParam := c.Query("latest_time"); latestTimeParam != "" {
	//	var err error
	//	latestTime, err = strconv.ParseInt(latestTimeParam, 10, 64)
	//	if err != nil {
	//		c.JSON(http.StatusOK, Response{
	//			StatusCode: CodeFailed,
	//			StatusMsg:  ParamsError,
	//		})
	//		return
	//	}
	//} else {
	//	// 没有接收到latest_time参数，设置为当前时间
	//	latestTime = time.Now().Unix()
	//}
	//
	//token := c.Query("token")
	//// 根据token获取用户id
	//userId, _ := GetIdByToken(token)
	//// 返回视频给用户
	//if videoLists := service.Feed(userId, latestTime); len(videoLists) > 0 {
	//	c.JSON(http.StatusOK, FeedResponse{
	//		Response:  Response{StatusCode: 0},
	//		VideoList: videoLists,
	//		NextTime:  time.Now().Unix(),
	//	})
	//}

}

//func GetIdByToken(token string) (int64, error) {
//
//}

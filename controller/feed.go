package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []entity.ApiVideo `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: nil,
		NextTime:  time.Now().Unix(),
	})
}

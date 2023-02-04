package entity

import "github.com/RaymondCode/simple-demo/controller"

type FeedResponse struct {
	controller.Response
	VideoList []VideoInfo `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

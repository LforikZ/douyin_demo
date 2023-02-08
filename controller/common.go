package controller

import "github.com/RaymondCode/simple-demo/entity"

var (
	CodeSuccess int32 = 0
	CodeFailed  int32 = 1
)

const (
	UploadSuccess = " uploaded successfully"
	UserNotExit   = "User doesn't exist"
	ParamsError   = "Params  error"
	VideosNotExit = "Videos doesn't exist"
	StatusSuccess = "Success"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []entity.ApiVideo `json:"video_list"`
}

type FeedResponse struct {
	Response
	VideoList []entity.ApiVideo `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

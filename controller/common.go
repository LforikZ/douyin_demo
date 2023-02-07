package controller

import "github.com/RaymondCode/simple-demo/entity"

var (
	CodeSuccess int32 = 0
	CodeFailed  int32 = 1
)

const (
	UploadSuccess        = "uploaded successfully"
	InsertCommentSuccess = "Insert comment successfully"
	DeleteCommentSuccess = "Delete comment successfully"
	UserNotExit          = "User doesn't exist"
	ParamsError          = "Params  error"
	ValidatorError       = "Validator Error"
	VideosNotExit        = "Videos doesn't exist"
	InsertCommentError   = "Insert comment Error"
	DeleteCommentError   = "Delete comment Error"
	ContentLoss          = "Content loss"
	CommentIDLoss        = "CommentID loss"
	GetCommentListFiled  = "Get CommentList Filed"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ResponseValim struct {
	Response
	Data interface{} `json:"data"`
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

type CommentResponse struct {
	Response
	Comment entity.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment entity.Comment `json:"comment,omitempty"`
}

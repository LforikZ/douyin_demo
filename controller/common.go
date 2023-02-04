package controller

var (
	CodeSuccess int32 = 0
	CodeFailed  int32 = 1
)

const (
	UploadSuccess = " uploaded successfully"
	UserNotExit   = "User doesn't exist"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//user
var (
	ErrorExistUser      int32 = 10002
	ErrorNotExistUser   int32 = 10003
	ErrorFailEncryption int32 = 10006
	ErrorNotCompare     int32 = 10007

	ErrorAuthCheckTokenFail    int32 = 30001 //token 错误
	ErrorAuthCheckTokenTimeout int32 = 30002 //token 过期
	ErrorAuthToken             int32 = 30003
	ErrorAuth                  int32 = 30004
	ErrorDatabase              int32 = 40001
)

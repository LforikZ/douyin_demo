package e

var (
	CodeSuccess int32 = 0
	CodeFailed  int32 = 1
)

const (
	UploadSuccess = " uploaded successfully"
	UserNotExit   = "User doesn't exist"
	//user
	RegisterSuccess     = "Register successfully"
	ErrorExistUser      = "User already exists"
	ErrorNotExistUser   = "User not exits"
	ErrorParams         = "Params Error"
	ErrorFailEncryption = "Password encryption failed"
	UserSelectSuccess   = "user select successfully"
	LoginSuccess        = "Login successfully"
	ErrorUserNotFound   = "User not found"

	ErrorNotCompare = "Password error"

	ErrorNotToken              = "Token not found"
	ErrorAuthCheckTokenFail    = "Token authentication failed" //token 错误
	ErrorAuthCheckTokenTimeout = "Token timed out"             //token 过期
	ErrorAuthToken             = "Token generation failed"
	ErrorAuth                  = "Token mismatch"
	ErrorDatabase              = "Database operation error, please try again"
	FindUserSuccess            = "Find user successfully"
	//follow
	ErrorInvalidOperation = "invalid operation"
	FollowSuccess         = "Follow successfully"
	UnfollowSuccess       = "Unfollow successfully"
	Cannotfollowyourself  = "Can not follow yourself"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

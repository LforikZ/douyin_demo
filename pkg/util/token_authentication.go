package util

import (
	"github.com/RaymondCode/simple-demo/pkg/e"
	"time"
)

//验证token
//传入参数：token
//输出：是否通过验证，错误信息。如果验证通过，错误信息为空。
func Authentication(token string) (bool, string) {
	res := false
	msg := ""

	if token == "" {
		msg = e.ErrorNotToken
		return res, msg
	} else {
		claims, err := ParseToken(token)
		if err != nil {
			msg = e.ErrorAuthCheckTokenFail
			return res, msg
		} else if time.Now().Unix() > claims.ExpiresAt {
			msg = e.ErrorAuthCheckTokenTimeout
			return res, msg
		}
	}
	res = true
	return res, msg
}

// @Author Junyi Tan 2023/2/3 23:30:00
package util

import (
	"errors"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/e"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id            int64  `json:"id"`
	Username      string `json:"username"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count""`
	IsFollow      bool   `json:"is_follow"`
	Authority     int64  `json:"authority"`
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(id int64, username string, follow_count int64,
	follower_count int64, is_follow bool, authority int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:            id,
		Username:      username,
		FollowCount:   follow_count,
		FollowerCount: follower_count,
		IsFollow:      is_follow,
		Authority:     authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "to-do-list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

//通过token获取用户信息
func GetUserInfo(token string) (*entity.User, error) {
	var user entity.User
	var claims *Claims
	auth, msg := Authentication(token)
	if auth == false {
		return nil, errors.New(msg)
	}
	claims, err := ParseToken(token)
	user.Id = claims.Id
	user.Name = claims.Username
	user.FollowCount = claims.FollowCount
	user.FollowerCount = claims.FollowerCount
	user.IsFollow = claims.IsFollow
	return &entity.User{Name: user.Name, Id: int64(user.Id),
		FollowCount: user.FollowCount, FollowerCount: user.FollowerCount,
		IsFollow: user.IsFollow}, err
}

func GetUserId(token string) int64 {
	user, _ := GetUserInfo(token)
	return user.Id
}

func GetUserName(token string) string {
	user, _ := GetUserInfo(token)
	return user.Name
}

func GetUserFollowCount(token string) int64 {
	user, _ := GetUserInfo(token)
	return user.FollowCount
}

func GetUserFollowerCount(token string) int64 {
	user, _ := GetUserInfo(token)
	return user.FollowerCount
}

func GetUserIsFollow(token string) bool {
	user, _ := GetUserInfo(token)
	return user.IsFollow
}

// Authentication 验证token
// 传入参数：token
// 输出：是否通过验证，错误信息。如果验证通过，错误信息为空。
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

package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]entity.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []entity.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	// 1.参数校验
	p := new(entity.ParamAction)
	if err := c.ShouldBindQuery(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			c.JSON(http.StatusOK, Response{
				StatusCode: CodeFailed,
				StatusMsg:  ParamsError,
			})
			return
		}
		errData := RemoveTopStruct(errs.Translate(Trans)) // 翻译并去除错误中结构体名字
		c.JSON(http.StatusOK, ResponseValim{
			Response: Response{
				StatusCode: CodeFailed,
				StatusMsg:  ValidatorError,
			},
			Data: errData,
		})
		return
	}

	// 2..验证token
	authentication, s := util.Authentication(p.Token)
	if authentication == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: CodeFailed,
			StatusMsg:  s,
		})
		return
	}
	// 3.验证发送用户是否存在
	id := util.GetUserId(p.Token)
	if id != 0 {
		userIdB, _ := strconv.Atoi(p.ToUserID)
		chatKey := genChatKey(id, int64(userIdB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := entity.Message{
			Id:         messageIdSequence,
			Content:    p.Content,
			CreateTime: time.Now().Format(time.Kitchen),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []entity.Message{curMessage}
		}
		c.JSON(http.StatusOK, Response{StatusCode: CodeSuccess})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: CodeFailed, StatusMsg: UserNotExit})
	}

}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}

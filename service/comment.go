// @Author Zihao_Li 2023/2/7 13:03:00
package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"time"
)

func InsertComment(p *entity.ParamComment) (err error) {
	typeOf := p.ActionType

	if typeOf == 1 {
		// 1-发布评论
		user, err := util.GetUserInfo(p.Token)
		if err != nil {
			return err
		}
		comment := &entity.Comment{
			CommentID:  util.GenID(),
			VideoID:    p.VideoID,
			User:       *user,
			Content:    p.Content,
			CreateDate: time.Now().String(),
		}
		err = mysql.InsertComment(comment)
		if err != nil {
			return err
		}
	} else {
		// 2-删除评论
		err := mysql.DeleteComment(p.CommentID)
		if err != nil {
			return err
		}
	}
	return nil
}

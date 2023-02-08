// @Author Zihao_Li 2023/2/7 13:03:00
package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
	"github.com/RaymondCode/simple-demo/pkg/util"
	"time"
)

func InsertComment(p *entity.ParamComment) (*entity.Comment, error) {
	typeOf := p.ActionType
	if typeOf == 1 {
		// 1-发布评论
		user, err := util.GetUserInfo(p.Token)
		if err != nil {
			return nil, err
		}
		comment := &entity.Comment{
			CommentID:  util.GetID(),
			VideoID:    p.VideoID,
			User:       *user,
			Content:    p.Content,
			CreateDate: time.Now().Format("01-02"),
		}
		err = mysql.InsertComment(comment)
		if err != nil {
			return comment, err
		}
		return comment, err
	} else {
		// 2-删除评论
		err := mysql.DeleteComment(p.CommentID)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}

func GetCommentList(p *entity.ParamTokenVID) (comments []entity.Comment, err error) {
	result, err := mysql.GetVideoListByVid(p.VideoID)
	if err != nil {
		return nil, err
	}
	for _, comment := range result {
		// 获取用户信息
		user, err := mysql.GetUserInfo(int(comment.User.Id))
		if err != nil {
			err = errors.New("user not exit")
			return nil, err
		}
		// 数据整合
		midData := entity.Comment{
			CommentID:  comment.CommentID,
			VideoID:    comment.VideoID,
			User:       *user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		comments = append(comments, midData)
	}
	return comments, err
}

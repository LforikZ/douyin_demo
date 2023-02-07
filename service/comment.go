// @Author Zihao_Li 2023/2/7 13:03:00
package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mysql"
)

func InsertComment(p *entity.ParamComment) {
	typeOf := p.ActionType
	if typeOf == 1 {
		mysql.InsertComment()
	} else {

	}
}

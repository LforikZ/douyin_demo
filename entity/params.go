// @Author Zihao_Li 2023/2/4 17:24:00
package entity

type ParamTokenUID struct {
	Token  string `form:"token" binding:"required"`
	UserID string `form:"user_id" binding:"required"`
}

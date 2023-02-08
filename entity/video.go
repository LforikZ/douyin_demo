// @Author Zihao_Li 2023/2/3 14:52:00
package entity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorID      int    //作者id
	AuthorName    string //作者姓名
	PlayUrl       string //视频地址
	CoverUrl      string //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
	Title         string //视频标题
}

type ApiVideo struct {
	Id            int64  //视频id，唯一标识
	*User                //嵌入作者信息
	PlayUrl       string //视频地址
	CoverUrl      string //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
	Title         string //视频标题
}

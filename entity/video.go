// @Author Zihao_Li 2023/2/3 14:52:00
package entity

type Video struct {
	AuthorID      string //作者id
	AuthorName    string //作者姓名
	PlayUrl       string //视频地址
	CoverUrl      string //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
}

type ApiVideo struct {
	*User                //嵌入用户信息
	PlayUrl       string //视频地址
	CoverUrl      string //封面网址
	FavoriteCount int64  //收藏计数
	CommentCount  int64  //评论计数
	IsFavorite    bool   //是否收藏
}

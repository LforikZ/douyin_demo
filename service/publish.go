// @Author Zihao_Li 2023/2/3 14:11:00
package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"mime/multipart"
	"path/filepath"
)

func Publish(data *multipart.FileHeader, user entity.User) (string, string) {

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	return saveFile, filename

}

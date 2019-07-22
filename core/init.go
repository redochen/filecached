package core

import (
	"github.com/redochen/filecached/models"
)

var (
	Files []*models.FileInfo = nil //文件列表
)

//初始化文件列表
func init() {
	Files = make([]*models.FileInfo, 0)
}

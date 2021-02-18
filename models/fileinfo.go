package models

import (
	"time"
)

//FileInfo 文件信息类
type FileInfo struct {
	Path string    //文件路径
	Time time.Time //创建时间
}

package models

import (
	"time"
)

type FileInfo struct {
	Path string    //文件路径
	Time time.Time //创建时间
}

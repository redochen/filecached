package models

type Directory struct {
	Category   string `json:"category"`   //分类名称
	Path       string `json:"path"`       //路径设置
	RecycleBin string `json:"recycleBin"` //回收站设置（为空表示过期直接删除文件）
	Duration   int    `json:"duration"`   //生命周期（单位：分钟，<=0 表示持久化存储）
}

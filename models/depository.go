package models

import (
	"encoding/json"
	"fmt"
	. "github.com/redochen/tools/string"
)

//路径设置
type Depository struct {
	Directories map[string]*Directory //目录设置
}

//获取目录
func (d *Depository) GetDirectory(category string) *Directory {
	if len(category) == 0 {
		return nil
	}

	if nil == d.Directories || len(d.Directories) == 0 {
		return nil
	}

	dir, _ := d.Directories[category]
	return dir
}

//设置目录
func (d *Depository) SetDirectory(category string, directory *Directory) {
	if len(category) == 0 || nil == directory {
		return
	}

	if nil == d.Directories {
		d.Directories = make(map[string]*Directory)
	}

	d.Directories[category] = directory
}

//解析目录设置
func (d *Depository) Parse(setting string) bool {
	if len(setting) == 0 {
		fmt.Println("[Parse] setting is empty.")
		return false
	}

	var directories []*Directory
	json.Unmarshal([]byte(setting), &directories)

	if nil == directories || len(directories) == 0 {
		fmt.Println("[Parse] json.Unmarshal failed.")
		return false
	}

	for _, item := range directories {
		d.SetDirectory(item.Category, item)
	}

	return true
}

//设置目录
func (d *Depository) Set(name string, section map[string]string) {
	if len(name) <= 0 || nil == section {
		return
	}

	dir := new(Directory)
	dir.Category = name
	dir.Path = section["path"]
	dir.RecycleBin = section["recycleBin"]
	dir.Duration = CcStr.ParseInt(section["duration"])
	//dir.Interval = HxStr.ParseInt(section["interval"])

	d.SetDirectory(dir.Category, dir)
}

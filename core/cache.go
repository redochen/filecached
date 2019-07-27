package core

import (
	cfg "github.com/redochen/filecached/config"
	. "github.com/redochen/filecached/models"
	"github.com/redochen/tools/file"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/log"
	"os"
)

//设置缓存
func SetCache(category, filename string, data []byte) bool {
	defer CheckPanic()

	d := getDirectory(category)
	if nil == d {
		return false
	}

	p := d.Path + "/" + filename

	_, err := os.Stat(p)
	if nil == err || os.IsExist(err) {
		os.Remove(p)
	}

	f, err := file.Open(p, true, false)
	if err != nil {
		Logger.Errorf("file.Open 《%s》 error: %s", p, err.Error())
		return false
	}

	defer f.Close()

	_, err = f.WriteEx(data, 0, cfg.UseGzip)
	if err != nil {
		Logger.Errorf("fe.WriteEx 《%s》 error: %s", p, err.Error())
		return false
	}

	return true
}

//获取缓存
func GetCache(category, filename string) []byte {
	defer CheckPanic()

	d := getDirectory(category)
	if nil == d {
		return nil
	}

	p := d.Path + "/" + filename

	f, err := file.Open(p, false, true)
	if err != nil {
		Logger.Errorf("file.Open 《%s》 error: %s", p, err.Error())
		return nil
	}

	defer f.Close()

	len, err := f.Size()
	if len <= 0 || err != nil {
		Logger.Errorf("fe.Size 《%s》 error: %s", p, err.Error())
		return nil
	}

	data := make([]byte, len)

	_, err = f.ReadEx(data, 0, cfg.UseGzip)
	if err != nil {
		Logger.Errorf("fe.ReadEx 《%s》 error: %s", p, err.Error())
	}

	return data
}

//获取目录
func getDirectory(category string) *Directory {
	dir := cfg.Depository.GetDirectory(category)
	if nil == dir {
		return cfg.Default
	} else {
		return dir
	}
}

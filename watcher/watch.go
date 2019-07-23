package main

import (
	"errors"
	. "github.com/redochen/filecached/models"
	. "github.com/redochen/filecached/util"
	. "github.com/redochen/tools/log"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

//监视目录
func WatchDirectory(directory *Directory) {
	defer func() {
		if err := recover(); err != nil {
			Logger.ErrorEx("WatchDirectory", "%v; stack: %s", err, debug.Stack())
		}
	}()

	if nil == directory || len(directory.Path) == 0 || directory.Duration <= 0 {
		return
	}

	for {
		checkDirectory(directory)
		time.Sleep(1 * time.Minute)
	}
}

//检查目录
func checkDirectory(directory *Directory) {
	if nil == directory ||
		len(directory.Path) == 0 ||
		directory.Duration <= 0 {
		return
	}

	err := filepath.Walk(directory.Path, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		} else if fi.IsDir() {
			return nil
		} else if isExpired(fi.ModTime(), directory.Duration) {
			return removeFile(fi.Name(), path, directory.RecycleBin)
		}
		return nil
	})

	if err != nil {
		Logger.ErrorEx("checkDirectory", "filepath.Walk error: %s", err.Error())
	}
}

//判断是否过期
func isExpired(t time.Time, duration int) bool {
	expiry := t.Add(time.Duration(duration) * time.Minute).UTC()
	return time.Now().UTC().After(expiry)
}

//删除文件：未设置回收站目录，则直接删除
func removeFile(name, path, recycleBin string) error {
	if len(name) == 0 ||
		len(path) == 0 {
		return errors.New("name or path is nil")
	}

	if len(recycleBin) > 0 {
		return MoveFileToRecycleBin(name, path, recycleBin)
	} else {
		return DeleteFilePhysically(name, path)
	}
}

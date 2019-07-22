package util

import (
	"errors"
	"fmt"
	"github.com/redochen/tools/log"
	"os"
	"time"
)

//创建目录
func CreateDirectory(directory string) error {
	if len(directory) <= 0 {
		return errors.New("[CreateDirectory] parameter is nil")
	}

	_, err := os.Stat(directory)
	if err != nil {
		return os.Mkdir(directory, os.ModeDir)
	}

	return nil
}

//物理删除文件
func DeleteFilePhysically(name, path string) error {
	if len(path) == 0 {
		return errors.New("[DeleteFilePhysically] path can not be empty")
	}

	err := os.Remove(path)
	ts := time.Now().Local().Format("01-02 03:04:05")

	if err != nil {
		l := fmt.Sprintf("[DeleteFilePhysically] os.Remove《%s》error: ", name)
		log.Logger.Error(l, err.Error())
		return err
	} else {
		fmt.Printf("[%s]《%s》has been deleted.\n", ts, name)
		return nil
	}
}

//移到文件到回收站
func MoveFileToRecycleBin(name, path, recycleBin string) error {
	if len(recycleBin) == 0 {
		return errors.New("[MoveFileToRecycleBin] recycleBin can not be empty")
	}

	p := recycleBin + "/" + name
	err := os.Rename(path, p)
	ts := time.Now().Local().Format("01-02 03:04:05")

	if err != nil {
		l := fmt.Sprintf("[MoveFileToRecyclebin] os.Rename《%s》error: ", name)
		log.Logger.Error(l, err.Error())
		return err
	} else {
		fmt.Printf("[%s]《%s》has been removed.\n", ts, name)
		return nil
	}
}

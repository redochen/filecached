package util

import (
	"errors"
	. "github.com/redochen/tools/log"
	"os"
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
	//ts := GetNowString("MM-dd hh:mm:ss", false)

	if err != nil {
		Logger.Errorf("os.Remove《%s》error: %s", name, err.Error())
		return err
	} else {
		//Logger.DebugEx("[%s]《%s》has been deleted.\n", ts, name)
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
	//ts := GetNowString("MM-dd hh:mm:ss", false)

	if err != nil {
		Logger.Errorf("os.Rename《%s》error: %s", name, err.Error())
		return err
	} else {
		//Logger.DebugEx("[%s]《%s》has been removed.\n", ts, name)
		return nil
	}
}

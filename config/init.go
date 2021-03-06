package config

import (
	"fmt"
	"runtime"

	"github.com/Unknwon/goconfig"
	"github.com/redochen/filecached/models"
	"github.com/redochen/filecached/util"
	CcStr "github.com/redochen/tools/string"
)

var (
	// Port 端口号
	Port = 16198

	//Depository 仓库实例
	Depository = &models.Depository{}

	// Default 默认仓库
	Default *models.Directory = nil

	//UseGzip 是否启用GZip
	UseGzip = false
)

//初始化配置
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	cfg, err := goconfig.LoadConfigFile("app.conf")
	if err != nil {
		fmt.Printf("[init] failed to load config file: %s\n", err.Error())
		panic(err)
	}

	sec, err := cfg.GetSection(goconfig.DEFAULT_SECTION)
	if err != nil {
		fmt.Println("[init] failed to get default section")
		panic(err)
	}

	if sec["port"] == "" {
		fmt.Printf("[init] failed to read port setting: %s\n", err.Error())
		panic(err)
	} else {
		Port = CcStr.ParseInt(sec["port"])
	}

	UseGzip = CcStr.ParseBool(sec["gzip"])
	if UseGzip {
		fmt.Println("启用压缩：是")
	} else {
		fmt.Println("启用压缩：否")
	}

	sections := cfg.GetSectionList()
	for _, name := range sections {
		if name == goconfig.DEFAULT_SECTION || name == "log" {
			continue
		}

		section, err := cfg.GetSection(name)
		if err != nil {
			panic("failed to get section [" + name + "]")
		}

		Depository.Set(name, section)
	}

	if Depository != nil && Depository.Directories != nil {
		for _, v := range Depository.Directories {
			if nil == v || len(v.Path) == 0 {
				continue
			}

			util.CreateDirectory(v.Path)
			util.CreateDirectory(v.RecycleBin)
		}
	}

	Default = Depository.GetDirectory("default")
}

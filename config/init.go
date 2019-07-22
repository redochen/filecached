package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/redochen/filecached/models"
	. "github.com/redochen/filecached/util"
	. "github.com/redochen/tools/string"
	"runtime"
)

var (
	Port                         = 16198
	Depository                   = &models.Depository{}
	Default    *models.Directory = nil   //默认仓库
	UseGzip                      = false //是否启用GZip
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

			CreateDirectory(v.Path)
			CreateDirectory(v.RecycleBin)
		}
	}

	Default = Depository.GetDirectory("default")
}

package conf

import (
	"errors"
	"io/ioutil"
	"strings"
	// "fmt"

	"shopping/utils/conf/parsers"
)

type Conf struct {
	ConfMap map[string]ConfParser
}

type ConfParser interface {
	Load(string) error
	Reload() error
	All() interface{}
	String(string) string
}

var Instance   *Conf
var ConfPath   string = "./conf"
var ConfSuffix string = ".conf"

// 注册配置模块
func Register() (err error) {
	conf_files, err := ioutil.ReadDir(ConfPath)

	if err != nil {
		return
	}

	Instance = &Conf{
		ConfMap: make(map[string]ConfParser),
	}

	s := make(map[string]string)

	for _, file := range conf_files {
		if !file.IsDir() {
			filename           := file.Name()
			filepath           := ConfPath+"/"+filename
			filenameTrimSuffix := strings.TrimSuffix(filename, ConfSuffix) 

			s[filenameTrimSuffix] = filepath
		}
	}

	for name := range s {
		// TODO 工厂模式：这里switch判定是哪个parser
		parser := &parsers.IniParser{}
		
		if err = parser.Load(s[name]); err != nil {
			return
		}

		Instance.ConfMap[name] = parser
	}

	return
}

// 获取配置文件信息
func New(filename string) (parser ConfParser, err error) {
	if _, ok := Instance.ConfMap[filename]; !ok {
		// 这里不应该做重载的蠢设置，配置文件应该在进程运行时就应该配置好，应该单写一个接口来提供热重载
		err = errors.New("未找到该配置文件，请确认信息是否正确")
		return
	}

	parser = Instance.ConfMap[filename]
	return
}

// 重载配置文件
func Reload(filename string) {
	// 
}

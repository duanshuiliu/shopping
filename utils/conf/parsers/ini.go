package parsers

import (
	"gopkg.in/ini.v1"
	"strings"
)

type IniParser struct {
	reader *ini.File	
}

func (this *IniParser) Load(config_file string) (err error) {
	conf, err := ini.Load(config_file)

	if err != nil {
		this.reader = nil
		return
	}

	this.reader = conf
	return 
}

func (this *IniParser) Reload() (err error) {
	return this.reader.Reload()
}

func (this *IniParser) All() interface{} {
	sections := this.reader.Sections()

	result := make(map[string]map[string]string)

	for _, section := range sections {
		result[section.Name()] = section.KeysHash()
	}
	
	return result
} 

func (this *IniParser) String(key string) (value string) {
	value = ""

	if this.reader == nil {
		return
	}

	s          := strings.Split(key, "::")
	sectionStr := ""
	keyStr     := ""

	switch (len(s)) {
		case 1:
			keyStr = s[0]
		case 2:
			sectionStr = s[0]
			keyStr     = s[1]
		default:
			
	}
	
	iniSection := this.reader.Section(sectionStr)
	
	if iniSection == nil {
		return
	}

	value = iniSection.Key(keyStr).String()	
	return
}

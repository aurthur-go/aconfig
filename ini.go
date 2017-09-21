package aconfig

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Ini : 定义ini配置结构
type Ini struct {
	filepath string
	conflist []map[string]map[string]string
}

// SetIni : 初始化
func SetIni(filepath string) *Ini {
	ini := new(Ini)
	ini.filepath = filepath

	return ini
}

// GetSection : 获取某配置
func (ini *Ini) GetSection(section string) map[string]string {
	ini.ReadList()
	conf := ini.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == section {
				return value
			}
		}
	}
	return nil
}

// GetValue : 获取单项配置值
func (ini *Ini) GetValue(section, name string) string {
	ini.ReadList()
	conf := ini.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	return ""
}

// SetValue : 添加配置值
func (ini *Ini) SetValue(section, key, value string) bool {
	ini.ReadList()
	data := ini.conflist
	var ok bool
	var index = make(map[int]bool)
	var conf = make(map[string]map[string]string)
	for i, v := range data {
		_, ok = v[section]
		index[i] = ok
	}

	i, ok := func(m map[int]bool) (i int, v bool) {
		for i, v := range m {
			if v == true {
				return i, true
			}
		}
		return 0, false
	}(index)

	if ok {
		ini.conflist[i][section][key] = value
	} else {
		conf[section] = make(map[string]string)
		conf[section][key] = value
		ini.conflist = append(ini.conflist, conf)
	}
	return true
}

// DeleteValue : 删除配置值
func (ini *Ini) DeleteValue(section, name string) bool {
	ini.ReadList()
	data := ini.conflist
	for i, v := range data {
		for key := range v {
			if key == section {
				delete(ini.conflist[i][key], name)
				return true
			}
		}
	}
	return false
}

// ReadList : 读取配置
func (ini *Ini) ReadList() []map[string]map[string]string {

	file, err := os.Open(ini.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#":
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
			if ini.uniquappend(section) == true {
				ini.conflist = append(ini.conflist, data)
			}
		}

	}

	return ini.conflist
}

// CheckErr : 检查错误
func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

// uniquappend:
func (ini *Ini) uniquappend(conf string) bool {
	for _, v := range ini.conflist {
		for k := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}

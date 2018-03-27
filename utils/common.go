package utils

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

// 判断文件/文件夹是否存在
func IsExit(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// 删除下划线,首字母大写
func UnderlineToUperCase(str string) string {
	strs := strings.Split(str, "_")
	s := ""
	for _, v := range strs {
		r := []rune(v)
		if r[0] >= 97 && r[0] <= 122 {
			r[0] = r[0] - 32
		}
		s += string(r)
	}
	return s
}

// 根据方法名调用方法 测试
// funcs := map[string]interface{}{"foo": foo, "bar": bar}
// call(funcs, "foo")
// call(funcs, "bar", 1, 2, 3)
func Call(m map[string]interface{}, funcName string, params ...interface{}) (err error) {
	f := reflect.ValueOf(m[funcName])
	if f.Type().Kind().String() == "func" {
		num := f.Type().NumIn()
		Info(num, len(params))
		if num != len(params) {
			return errors.New("参数数量不匹配")
		}
		in := make([]reflect.Value, num)
		for i, v := range params {
			in[i] = reflect.ValueOf(v)
		}
		f.Call(in)
	} else {
		return errors.New("不是方法,调用失败")
	}
	return
}

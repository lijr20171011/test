package utils

import (
	"os"
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

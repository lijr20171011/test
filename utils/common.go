package utils

import (
	"errors"
	"os"
	"reflect"
	"sort"
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

// ==========切片排序 start================
// 仅限切片类型
func MySort(in interface{}, columnName string, orderType ...interface{}) (interface{}, error) {
	inType := reflect.TypeOf(in)
	inValue := reflect.ValueOf(in)
	var sortSlice mySortSlice // 排序切片
	// 判断是否为切片
	if inType.Kind().String() != "slice" {
		return nil, errors.New("非切片类型,排序异常")
	}
	// 获取切片长度
	inlen := inValue.Len()
	// 整理排序切片
	for i := 0; i < inlen; i++ {
		// 获取切片第i个元素信息
		v := inValue.Index(i)
		// 判断是否有指定字段
		_, ok := v.Type().FieldByName(columnName)
		if !ok {
			return nil, errors.New("获取指定字段异常")
		}
		columnInfo := sortColumnInfo{
			Index:       i,
			ColumnValue: v.FieldByName(columnName).Interface(), //指定字段的值
		}
		sortSlice = append(sortSlice, columnInfo)
	}
	// 排序
	sort.Sort(sortSlice)
	// 排序后切片信息
	re := reflect.MakeSlice(inType, inlen, inValue.Cap())
	// 获取排序方式 默认:正序
	order := true
	if len(orderType) > 0 && (orderType[0] == -1 || orderType[0] == "-1") {
		order = false
	}
	if order { // 正序
		for i, v := range sortSlice {
			re.Index(i).Set(inValue.Index(v.Index))
		}
	} else { // 倒序
		for i, v := range sortSlice {
			re.Index(inlen - i - 1).Set(inValue.Index(v.Index))
		}
	}
	return re.Interface(), nil
}

type sortColumnInfo struct {
	Index       int
	ColumnValue interface{}
}

type mySortSlice []sortColumnInfo

func (c mySortSlice) Len() int {
	return len(c)
}

func (c mySortSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c mySortSlice) Less(i, j int) bool {
	iType := reflect.TypeOf(c[i].ColumnValue).Kind().String()
	jType := reflect.TypeOf(c[j].ColumnValue).Kind().String()
	if iType == jType {
		// todo
		switch iType {
		case "string":
			return c[i].ColumnValue.(string) < c[j].ColumnValue.(string)
		case "int":
			return c[i].ColumnValue.(int) < c[j].ColumnValue.(int)
		case "int8":
			return c[i].ColumnValue.(int8) < c[j].ColumnValue.(int8)
		case "int16":
			return c[i].ColumnValue.(int16) < c[j].ColumnValue.(int16)
		case "int32":
			return c[i].ColumnValue.(int32) < c[j].ColumnValue.(int32)
		case "int64":
			return c[i].ColumnValue.(int64) < c[j].ColumnValue.(int64)
		case "uint":
			return c[i].ColumnValue.(uint) < c[j].ColumnValue.(uint)
		case "uint8":
			return c[i].ColumnValue.(uint8) < c[j].ColumnValue.(uint8)
		case "uint16":
			return c[i].ColumnValue.(uint16) < c[j].ColumnValue.(uint16)
		case "uint32":
			return c[i].ColumnValue.(uint32) < c[j].ColumnValue.(uint32)
		case "uint64":
			return c[i].ColumnValue.(uint64) < c[j].ColumnValue.(uint64)
		case "float32":
			return c[i].ColumnValue.(float32) < c[j].ColumnValue.(float32)
		case "float64":
			return c[i].ColumnValue.(float64) < c[j].ColumnValue.(float64)
		}
		// Bool
		// Uintptr
		// Complex64
		// Complex128
		// Array
		// Chan
		// Func
		// Interface
		// Map
		// Ptr
		// Slice
		// Struct
		// UnsafePointer
	}
	return false
}

// ==========切片排序 end================

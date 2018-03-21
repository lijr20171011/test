package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/orm"

	"github.com/urfave/cli"
	// go get github.com/urfave/cli
	// git clone https://github.com/urfave/cli.git

	_ "github.com/go-sql-driver/mysql"
)

type DBTableInfo struct {
	ColumnName    string `json:"column_name"`    // 字段名
	DataType      string `json:"data_type"`      // 数据类型
	ColumnComment string `json:"column_comment"` // 备注
	ColumnKey     string `json:"column_key"`     // 索引
}

func main() {
	// sqltogo()
	sqltogo1()
}

func sqltogo1() {

}

func sqltogo() {
	app := cli.NewApp()
	app.Name = "sqltogo"
	app.Usage = "用途:根据库名表名将表结构转化为go结构体"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Value: "wr",
			Usage: "数据库库名",
		},
		cli.StringFlag{
			Name:  "table",
			Value: "users",
			Usage: "表名",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
			Usage: "数据库主机地址",
		},
		cli.StringFlag{
			Name:  "user",
			Value: "root",
			Usage: "用户名",
		},
		cli.StringFlag{
			Name:  "pwd",
			Value: "123456",
			Usage: "密码",
		},
		cli.BoolFlag{
			Name:  "writefile",
			Usage: "是否写入文件",
		},
	}
	app.Action = func(c *cli.Context) error {
		dbName := c.String("db")
		tableName := c.String("table")
		host := c.String("host")
		user := c.String("user")
		pwd := c.String("pwd")
		writeFile := c.Bool("writefile")
		GetDBTableStruct(host, user, pwd, dbName, tableName, writeFile)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetDBTableStruct(host, user, pwd, dbName, tableName string, writeFile bool) {
	// 连接数据库
	err := orm.RegisterDataBase("default", "mysql", user+":"+pwd+"@tcp("+host+":3306)/"+dbName+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 查询
	var infos []DBTableInfo
	sql := `SELECT column_name,data_type,column_comment,column_key FROM information_schema.columns WHERE table_schema = ? AND table_name = ? `
	_, err = orm.NewOrm().Raw(sql, dbName, tableName).QueryRows(&infos)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 整理数据
	hasTime := false
	structStr := "\n\ntype " + UnderlineToUperCase(tableName) + " struct {\n"
	for _, v := range infos {
		// 字段名
		structStr += "\t" + UnderlineToUperCase(v.ColumnName) + "\t"
		// 字段类型
		switch v.DataType {
		case "int", "tinyint":
			structStr += "int" + "\t"
		case "float", "double", "decimal":
			structStr += "float64" + "\t"
		case "date", "datetime", "time", "timestamp":
			structStr += "time.Time" + "\t"
			hasTime = true
		case "char", "varchar", "text", "longtext":
			structStr += "string" + "\t"
		case "bigint":
			structStr += "int64" + "\t"
		default:
			fmt.Println("数据类型不明 --> " + v.DataType)
			return
		}
		// 字段说明
		structStr += "`orm:"
		structStr += `"column(` + v.ColumnName + `)`
		if v.ColumnKey == "PRI" {
			structStr += ";pk"
		}
		structStr += `"` + "`\t"
		// 注释
		if v.ColumnComment != "" {
			structStr += "// " + v.ColumnComment
		}
		structStr += "\n"
	}
	structStr += "}\n"
	if writeFile {
		// 写入文件
		headInfo := "package " + tableName + "\n"
		if hasTime {
			headInfo += "\nimport(\n\t" + `"time"` + "\n)\n"
		}
		structToFile(tableName, headInfo+structStr)
	} else {
		fmt.Println(structStr)
	}
	return
}

func structToFile(tableName, fileInfo string) {
	path := "/"
	// 获取当前系统分隔符
	if os.IsPathSeparator('\\') {
		path = "\\"
	}
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断models是否已存在
	modelsPaht := dir + path + "models"
	if !IsExit(modelsPaht) {
		// 不存在创建models目录
		err = os.Mkdir(modelsPaht, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// 判断 table文件是否已存在
	filePath := modelsPaht + path + tableName + ".go"
	if IsExit(filePath) {
		fmt.Println(errors.New("models目录下已有" + tableName + "文件"))
		return
	}
	// 创建tableName文件
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fileInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

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

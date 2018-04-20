package main

import (
	"my_project/test/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestStruct struct {
	Name string
	Age  int
}

func main() {
	utils.Info("测试开始")
	updateTest()
	utils.Info("测试结束")
}

var mgoSession *mgo.Session

func updateTest() {
	utils.Info("更新测试开始")
	session := getSession()
	if session == nil {
		utils.Info("获取session异常")
		return
	}
	defer session.Close()
	test1 := session.DB("test").C("test1")
	// 字段必须全部小写,否则找不到(新加一个字段),更新相应字段的值和类型
	// update := map[string]interface{}{"age": 11}
	// err := test1.Update(bson.M{"name": "aaa"}, bson.M{"$set": update})
	err := test1.Update(bson.M{"name": "aaa"}, bson.M{"$set": bson.M{"age": "11"}})
	if utils.IsErr(err) {
		return
	}
	utils.Info("更新测试结束")
}

func insertTest() {
	utils.Info("添加开始")
	test := TestStruct{
		Name: "aaa",
		Age:  22,
	}
	session := getSession()
	if session == nil {
		utils.Info("获取session异常")
		return
	}
	defer session.Close()
	// 若没有自动创建,字段名全转为小写
	err := session.DB("test").C("test1").Insert(&test)
	if utils.IsErr(err) {
		return
	}
	utils.Info("添加结束")
}

func getSession() *mgo.Session {
	url := "127.0.0.1:27017"
	var err error
	if mgoSession == nil {
		mgoSession, err = mgo.Dial(url)
		if utils.IsErr(err) {
			return nil
		}
	} else if mgoSession != nil {
		err = mgoSession.Ping()
		if utils.IsErr(err) {
			mgoSession, err = mgo.Dial("xxx")
			if utils.IsErr(err) {
				return nil
			}
		}
	}
	return mgoSession.Clone()
}

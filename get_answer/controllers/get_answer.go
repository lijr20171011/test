package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"my_project/get_answer/models"
	"my_project/get_answer/utils"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/chenqinghe/baidu-ai-go-sdk/version/ocr"
)

func StartAnswer() {
	utils.Info("开始")
	for {
		var cmd string
		fmt.Print(">")
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			GetAnswer()
		case "2":
			os.Exit(1)
		}
	}
	utils.Info("结束")
}

func GetAnswer() {
	imgName := "./img/ceshi2.png"
	GetAnswerByPicture(imgName)
}

//根据图片获取答案 (.png)
func GetAnswerByPicture(imgName string) (err error) {
	// 获取图片有效区域
	cutPath, err := GetCutPicture(imgName)
	if err != nil {
		utils.Info(err)
		return
	}
	// ai识别文字
	res, err := BaiduAi(cutPath)
	if err != nil {
		utils.Info(err)
		return
	}
	// 整理问题答案信息
	q, a, err := GetQA(res)
	if err != nil {
		utils.Info(err)
		return
	}
	utils.Info(q)
	// 发起请求
	ch := make(chan models.BaiduSearchResp, len(a))
	for i, v := range a {
		go SearchByKeyWord(i+1, q, v, ch)
	}
	var answer models.BaiduSearchResp
	//获取答案
	var resp models.BaiduSearchResp
	for i := 0; i < len(a); i++ {
		resp = <-ch
		if resp.Num > answer.Num {
			answer.Sort = resp.Sort
			answer.Key = resp.Key
			answer.Num = resp.Num
		}
	}
	utils.Info("------------------------------------------------------")
	utils.Info("|")
	utils.Info("|    最终答案:"+answer.Key, "数量:"+strconv.Itoa(answer.Num), "排序:          ", answer.Sort)
	utils.Info("|")
	utils.Info("------------------------------------------------------")
	return
}

// 获取图片有效区域
func GetCutPicture(pictureName string) (cutPath string, err error) {
	// 获取截图名称
	index := strings.LastIndex(pictureName, ".png")
	if index == -1 {
		err = errors.New("图片格式异常")
		return
	}
	_, fileName := path.Split(pictureName)
	cutPath = utils.SavePath + "cut_" + fileName
	// 打开图片文件
	f1, err := os.OpenFile(pictureName, os.O_RDONLY, 0777)
	if err != nil {
		return
	}
	defer f1.Close()
	m, err := png.Decode(f1)
	if err != nil {
		return
	}
	imgrgb := m.(*image.NRGBA)
	// todo 自动识别裁剪范围
	effectiveImg := imgrgb.SubImage(image.Rect(utils.X1, utils.Y1, utils.X2, utils.Y2))
	cutFile, err := os.Create(cutPath)
	if err != nil {
		return
	}
	defer cutFile.Close()
	err = png.Encode(cutFile, effectiveImg)
	return
}

// 百度ai识别
func BaiduAi(imgName string) (data models.BaiduAiResp, err error) {
	file, err := os.Open(imgName)
	if err != nil {
		return
	}
	client := ocr.NewOCRClient(utils.BaiduAiAppKey, utils.BaiduAiSecretKey)
	res, err := client.GeneralRecognizeBasic(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return
	}
	return
}

// 整理问题答案
func GetQA(data models.BaiduAiResp) (q string, a []string, err error) {
	l := len(data.WordsResult)
	if l <= 3 {
		err = errors.New("获取题目信息异常")
		return
	}
	for _, v := range data.WordsResult[:l-3] {
		q += v.Words
	}
	for _, v := range data.WordsResult[l-3:] {
		a = append(a, v.Words)
	}
	return
}

// 根据关键词查询
func BaiduSearch(i int, q, key string, ch chan models.BaiduSearchResp) (err error) {
	u := fmt.Sprintf("http://www.baidu.com/s?wd=%s", url.QueryEscape(q+" "+key))
	resp, err := http.Get(u)
	if err != nil {
		utils.Info(err)
		return
	}
	defer resp.Body.Close()
	res, err := ioutils.ReadAll(resp.Body)
	if err != nil {
		utils.Info(err)
		return
	}
	s := string(res)
	index := strings.Index(s, "百度为您找到相关结果约")
	if index == -1 {
		err = errors.New("查询失败")
		return
	}
	s = s[index+33:]
	index = strings.Index(s, "个")
	s = s[:index]
	s = strings.Replace(s, ",", "", -1)
	num, err := strconv.Atoi(s)
	if err != nil {
		utils.Info(err)
		return
	}
	utils.Info(i, key+":"+s)
	resp := models.BaiduSearchResp{
		Sort: i,
		Key:  key,
		Num:  num,
	}
	ch <- resp
	return
}

func SearchByKeyWord(i int, q, key string, ch chan models.BaiduSearchResp) {
	err := BaiduSearch(i, q, key, ch)
	if err != nil {
		utils.Info(key+"---查询失败", err.Error())
	}
}

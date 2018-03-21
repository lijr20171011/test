package models

// 百度ai识别数据
type BaiduAiResp struct {
	LogId          int64 `json:"log_id"`
	WordsResultNum int   `json:"words_result_num"`
	WordsResult    []struct {
		Words string `json:"words"`
	} `json:"words_result"`
}

// 百度搜索结果
type BaiduSearchResp struct {
	Sort int    //答案排序
	Key  string //答案内容
	Num  int    //答案搜索数量
}

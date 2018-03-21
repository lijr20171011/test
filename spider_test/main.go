package main

// import (
// 	_ "spider_test/routers"

// 	"github.com/astaxie/beego"
// )

// func main() {
// 	if beego.BConfig.RunMode == "dev" {
// 		beego.BConfig.WebConfig.DirectoryIndex = true
// 		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
// 	}
// 	beego.Run()
// }

import (
	"fmt"
	"net/http"
	"net/url"
	_ "spider_test/routers"

	"github.com/astaxie/beego"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
)

type IndexPageProcesser struct {
}

func NewIndexPageProcesser() *IndexPageProcesser {
	return &IndexPageProcesser{}
}

func main() {

}

func (this *IndexPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}
	query := p.GetHtmlParser()
	h_url, _ := query.Find("#wrapper .item:nth-child(7) .item_title a").Attr("href")

	// t, _ := query.Find("#wrapper").Html()
	// beego.Info(t)

	// query.Find("#wrapper .item").Each(func(i int, s *goquery.Selection) {
	// 	beego.Info(i)
	// 	a, _ := s.Html()
	// 	beego.Info(a)
	// })

	beego.Info("=======")
	beego.Info(h_url)
	beego.Info("=======")
}

func RunSpider() {
	reqs := make([]*request.Request, 0)
	header := http.Header{}
	// header.Add("referer", "https://studygolang.com/")
	// header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	// header.Add("Accept-Encoding", "gzip, deflate, br")
	// header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	urlKey := "福"
	urlKey = url.QueryEscape(urlKey)
	cookes := make([]*http.Cookie, 0)
	cookie := &http.Cookie{}
	cookes = append(cookes, cookie)
	cookie.Value = "BDTUJIAID=ecda3b112ecaa47601f03c57f745ef73; UM_distinctid=15e2c1fb6324eb-06732e9bf5ee86-464c0328-1fa400-15e2c1fb63328e; _ga=GA1.2.292387503.1513228215; Hm_lvt_224c227cd9239761ec770bc8c1fb134c=1517899091; _gid=GA1.2.625365945.1517899091; user=MTUxNzg5OTQ4MHxEdi1CQkFFQ180SUFBUkFCRUFBQUp2LUNBQUVHYzNSeWFXNW5EQXNBQ1VsT1JFVllYMVJCUWdaemRISnBibWNNQlFBRFlXeHN8PWZY50F5cSZaYXbE1eGfl1lSzgEaSKKGQ5JZxT4ls5I=; Hm_lpvt_224c227cd9239761ec770bc8c1fb134c=1517899481"
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req := request.NewRequest("https://studygolang.com/", "html", "", "GET", "", header, cookes, nil, nil)
	reqs = append(reqs, req)
	spider.NewSpider(NewIndexPageProcesser(), "TaskName").
		AddRequests(reqs).
		SetThreadnum(1).
		Run()
}

func (this *IndexPageProcesser) Finish() {
	fmt.Printf("第一层路径爬取完毕")
}
func init() {
	RunSpider()
}

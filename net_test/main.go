package main

import (
	"fmt"
	"io/ioutil"
	"my_project/net_test/utils"
	"net/http"
	"time"
)

// func main() {
// 	utils.Info("程序开始")
// 	// netTest1()
// 	netTest2()
// 	// netTest3()
// 	utils.Info("程序结束")
// }

func netTest3() {
	go start()
	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8555/测试")
	utils.ExitWithErr(err)
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	utils.Info(string(r))
}

func netTest2() {
	// client := &http.Client{}
	// resp, err := client.Get("http://example.com")
	resp, err := http.NewRequest("GET", "http://example.com", nil)
	utils.ExitWithErr(err)
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	utils.ExitWithErr(err)
	utils.Info(string(r))

}

func netTest1() {
	go start()
	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8555/测试")
	utils.ExitWithErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	utils.ExitWithErr(err)
	utils.Info(string(body))
}

func start() {
	utils.Info("开始=========")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("127.0.0.1:8555", nil)
	utils.ExitWithErr(err)
	utils.Info("结束=========")
}

func handler(w http.ResponseWriter, r *http.Request) {
	utils.Info("==测试输出===", r.URL.Path)
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	utils.ExitWithErr(err)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	utils.Info("pos(", i, ") :", pos(i))
	// 	utils.Info("neg(-2*", i, ") :", neg(-2*i))
	// }
}

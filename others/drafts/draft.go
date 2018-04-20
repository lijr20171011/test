package main

import (
	"errors"
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"../../utils"
)

func main() {
	StringCompare()
}

func StringCompare() {
	utils.Info("12" > "3")
}

func PrintWithColorTest() {
	// color.Set(color.FgMagenta, color.Bold)
	// defer color.Unset()
	// fmt.Println("aaaa")
	// color.Set(color.FgGreen, color.Bold)
	// fmt.Println("bbb")
	// fmt.Printf("\033[1;31;40m%s", "aaaa")
	fmt.Printf("\x1b[%dm%s\x1b[0m", 91, "test")
}

func MaxCommonDivisor(x, y int) int {
	var diff int
	for {
		diff = x % y
		if diff > 0 {
			x = y
			y = diff
		} else {
			return y
		}
	}
}

var (
	t1 = time.Now().AddDate(0, 0, 1)
	t2 = time.Now().AddDate(0, 0, 1)
	t3 = time.Now().AddDate(0, 0, 1)
)

// func main() {
// 	time_test_1()
// 	time.Sleep(3 * time.Second)
// 	time_test_1()
// }

func time_test_1() {
	utils.Info(t1)
	time.Sleep(2 * time.Second)
	utils.Info(t2)
	time.Sleep(2 * time.Second)
	utils.Info(t3)
}

func math_pow_test() {
	a := 2.0
	b := 3.0
	c := math.Pow(a, b)
	utils.Info(c)
}

func errTest() (err error) {
	err = errors.New("ceshi")
	return errors.New("sss")
}

func sinceTime(start time.Time) {
	utils.Info(start)
	utils.Info(time.Now())
}

func timeTest() {
	t1 := time.Now().AddDate(0, -4, -5)
	utils.Info(t1)
	now := time.Now()
	utils.Info(now.Sub(t1))
}

func timeFormatTest() {
	formatStr := "20060102.api"
	timeStr := time.Now().Format(formatStr)
	utils.Info(timeStr)
}

// func main() {
// 	idCard := "330382199506080924"
// 	var reg *regexp.Regexp
// 	placeStr1 := "**"
// 	placeStr2 := ""
// 	var err error
// 	if len(idCard) == 15 {
// 		reg, err = regexp.Compile("^(\\d{4})(\\d{2})(\\d{2})(\\d{5})(.*)")
// 		placeStr2 = "*****"
// 	} else if len(idCard) == 18 {
// 		reg, err = regexp.Compile("^(\\d{4})(\\d{2})(\\d{4})(\\d{6})(.*)")
// 		placeStr2 = "******"
// 	} else {
// 		utils.Info(len(idCard))
// 		return
// 	}
// 	if err != nil {
// 		utils.Info(err)
// 		return
// 	}
// 	if reg.MatchString(idCard) == true {
// 		submatch := reg.FindStringSubmatch(idCard)
// 		utils.Info(len(submatch))
// 		utils.Info(submatch[0])
// 		utils.Info(submatch[1] + placeStr1 + submatch[3] + placeStr2 + submatch[5])
// 	}
// }

// func main() {
// 	// var m map[string]string
// 	m := make(map[string]string)
// 	utils.Info(m == nil)
// 	// s := []string{}
// 	var s []string
// 	utils.Info(s == nil)
// 	// sync_test2()
// }

func sync_test2() {

}

func a(ch chan int) {
	ch <- 1
}

func chan_test2() {
	i := 0
	ch := make(chan int)
	for i < 10 {
		go a(ch)
		<-ch
		i++
		utils.Info(i)
	}
	defer func() {
		utils.Info("===")
		go utils.Info("===")
	}()
}

// func main() {
// 	// f := new(myfunc)
// 	// r := new(RunTime)
// 	// utils.Info(r.RunTime)
// 	t := new(S)
// 	utils.Info(t)
// }

// type S struct {
// 	Atime time.Time
// }

// func start() time.Time {
// 	t := time.Now()
// 	return t
// }

// // 定义方法
// type myfunc interface {
// 	init(c)
// }

// type RunTime struct {
// 	RunTime time.Time
// }

// func (t RunTime) init() {
// 	t.RunTime = time.Now()
// 	utils.Info(t)
// }

// func start(t time.Time) {
// 	t = time.Now()
// }

// func end(t time.Time) {
// 	utils.Info(time.Since(t).String())
// }

func reflect_test() {
	funcs := map[string]interface{}{"foo": foo, "bar": bar}
	utils.Call(funcs, "foo")
	utils.Call(funcs, "bar", 1, 2, 3)
}

func foo() {
	utils.Info("foo")
}
func bar(a, b, c int) {
	utils.Info("bar")
}

// 关于获取方法执行时间
func time_test() {
	defer utils.GetFuncRunTime(time.Now())
	pointer_test()
}

// 关于字符串
func letter_test() {
	s := "中国人"
	fmt.Println("s byte sequence:")
	for i := 0; i < len(s); i++ {
		// 字节序
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Println("")
}

type field struct {
	name string
}

func print(p *field) {
	fmt.Println(p.name)
}

// 关于指针的使用 for range只有1个v
func pointer_test() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go print(v)
	}
	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go print(&v)
		runtime.Gosched()
	}
	time.Sleep(1 * time.Second)
}

// func (p *field) print() {
// 	fmt.Println(p.name)
// }

// func main() {
// 	data1 := []*field{{"one"}, {"two"}, {"three"}}
// 	for _, v := range data1 {
// 		go v.print()
// 	}
// 	time.Sleep(2 * time.Second)
// 	data2 := []field{{"four"}, {"five"}, {"six"}}
// 	for _, v := range data2 {
// 		go v.print()
// 	}

// 	time.Sleep(3 * time.Second)
// }

// chan测试
func chan_test1() {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	v, ok := <-c
	utils.Info(v, ok)
	close(c)
	v, ok = <-c
	utils.Info(v, ok)
}

// sync测试
func sync_test1() {
	var once sync.Once
	onceBody := func() {
		utils.Info("Only once.")
	}
	done := make(chan bool)
	utils.Info(len(done))
	utils.Info(cap(done))
	for i := 0; i < 10; i++ {
		go func() {
			utils.Info("o.done >> ", once)
			once.Do(onceBody)
			utils.Info("输出")
			done <- true
			utils.Info("len >>> ", len(done))
			utils.Info("cap >>> ", cap(done))
		}()
	}
	for i := 0; i < 10; i++ {
		tf := <-done
		utils.Info(tf)
		utils.Info(i)
	}
}

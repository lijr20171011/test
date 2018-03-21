package main

import (
	"fmt"
	"math"
	"my_project/get_answer/utils"
	"sort"
	"strconv"
	"unsafe"
)

// ======================== 百度AI识别
/*
	// ======================== 百度AI识别
	func main() {
		utils.Info("程序执行")
		controllers.StartAnswer()
		utils.Info("程序结束")
	}
*/

// ========================= 定时 例2:
/*
	// ========================= 定时 例2:
	// 发送者
	func sender(c chan int) {
		for i := 0; i < 100; i++ {
			c <- i
			if i >= 5 {
				time.Sleep(time.Second * 7)
			} else {
				time.Sleep(time.Second)
			}
		}
	}

	func main() {
		c := make(chan int)
		go sender(c)
		timeout := time.After(time.Second * 3)
		for {
			select {
			case d := <-c:
				fmt.Println(d)
			case <-timeout:
				fmt.Println("这是定时操作任务 >>>>>")
			case dd := <-time.After(time.Second * 3):
				fmt.Println(dd, "这是超时*****")
			}

			fmt.Println("for end")
		}
	}
*/

// ====================== 定时 例1:
/*
	// ====================== 定时 例1:
	func main() {
		//closeChannel()
		c := make(chan int)
		timeout := time.After(time.Second * 2) //
		t1 := time.NewTimer(time.Second * 3)   // 效果相同 只执行一次
		var i int
		go func() {
			for {
				select {
				case <-c:
					fmt.Println("channel sign")
					return
				case <-t1.C: // 代码段2
					fmt.Println("3s定时任务")
				case <-timeout: // 代码段1
					i++
					fmt.Println(i, "2s定时输出")
				case <-time.After(time.Second * 4): // 代码段3
					fmt.Println("4s timeout。。。。")
				default: // 代码段4
					fmt.Println("default")
					time.Sleep(time.Second * 1)
				}
			}
		}()
		time.Sleep(time.Second * 6)
		close(c)
		time.Sleep(time.Second * 2)
		fmt.Println("main退出")
	}
*/

// =============================== switch中跳出for循环测试
/*
	// =============================== switch中跳出for循环测试
	func init() {
		strs := []string{"a", "b", "c", "d"}
		strss := []string{"a", "b"}
		str := [][]string{strs, strss, strs}
		for i, v := range str {
			abd := "abd"
			abc := "abc"
			utils.Info(i)
			abd = "aaa"
		Loop:
			for ii, vv := range v {
				utils.Info(ii, vv)
				switch ii {
				case 2:
					utils.Info("===")
					abd = "bbb"
					break Loop
				default:
					abd = "vvv"
					utils.Info("-----")
				}
			}
			utils.Info(abc)
			utils.Info(abd)
		}
	}
*/

// golang select 用法详解 http://blog.csdn.net/wo198711203217/article/details/65442288

// func main() {
// 	var chs1 = make(chan int)
// 	var chs2 = make(chan float64)
// 	var chs3 = make(chan string)
// 	var ch4close = make(chan int)
// 	defer close(ch4close)

// 	go func(c chan int, ch4close chan int) {
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 		}
// 		close(c)
// 		ch4close <- 1
// 	}(chs1, ch4close)

// 	go func(c chan float64, ch4close chan int) {
// 		for i := 0; i < 5; i++ {
// 			c <- float64(i) + 0.1
// 		}
// 		close(c)
// 		ch4close <- 1
// 	}(chs2, ch4close)

// 	go func(c chan string, ch4close chan int) {
// 		for i := 0; i < 5; i++ {
// 			c <- "string:" + strconv.Itoa(i)
// 		}
// 		close(c)
// 		ch4close <- 1
// 	}(chs3, ch4close)

// 	var selectCase = make([]reflect.SelectCase, 4)
// 	selectCase[0].Dir = reflect.SelectRecv
// 	selectCase[0].Chan = reflect.ValueOf(chs1)

// 	selectCase[1].Dir = reflect.SelectRecv
// 	selectCase[1].Chan = reflect.ValueOf(chs2)

// 	selectCase[2].Dir = reflect.SelectRecv
// 	selectCase[2].Chan = reflect.ValueOf(chs3)

// 	selectCase[3].Dir = reflect.SelectRecv
// 	selectCase[3].Chan = reflect.ValueOf(ch4close)

// 	done := 0
// 	finished := 0
// 	for finished < len(selectCase)-1 {
// 		chosen, recv, recvOk := reflect.Select(selectCase)

// 		if recvOk {
// 			done = done + 1
// 			switch chosen {
// 			case 0:
// 				fmt.Println(chosen, recv.Int())
// 			case 1:
// 				fmt.Println(chosen, recv.Float())
// 			case 2:
// 				fmt.Println(chosen, recv.String())
// 			case 3:
// 				finished = finished + 1
// 				done = done - 1
// 				// fmt.Println("finished\t", finished)
// 			}
// 		}
// 	}
// 	fmt.Println("Done", done)

// }

// func main() {
// 	var s interface{}
// 	s = 0
// 	v := s.(int)
// 	// v := reflect.ValueOf(s).Int()
// 	utils.Info(v)
// }

// defer()测试
// func main() {
// 	utils.Info("开始")
// 	defer func() {
// 		utils.Info("=====")
// 	}()
// 	for i := 0; i < 5; i++ {
// 		flag := false
// 		utils.Info(i)
// 		defer func() {
// 			utils.Info(flag)
// 		}()
// 		if i == 3 {
// 			flag = true
// 		}
// 	}
// 	utils.Info("测试中")
// 	defer func() {
// 		utils.Info("进行时")
// 	}()
// 	utils.Info("结束")
// }

// func main() {
// 	str := "[([{])]}"
// 	utils.Info(">>>", isValid(str))
// }

var bracketsMap map[byte]byte = map[byte]byte{'(': '-', ')': '(', '[': '-', ']': '[', '{': '-', '}': '{'}

//判断括号是否正确关闭
func isValid(s string) bool {
	slen := len(s)
	brackets := make([]byte, 0, slen)
	for i := 0; i < slen; i++ {
		if m, ok := bracketsMap[s[i]]; ok { //括号
			//判断是否为左括号
			if m == '-' { //左括号
				brackets = append(brackets, s[i])
				continue
			}
			// 右括号
			l := len(brackets)
			if l == 0 || brackets[l-1] != m {
				return false
			}
			brackets = brackets[:l-1]
		}
	}
	if len(brackets) == 0 {
		return true
	}
	return false
}

// func main() {
// 	strs := []string{
// 		// "abce",
// 		// "abcdaf",
// 		// "abe",
// 		// "abfje",
// 		// "a",
// 		// "aa",
// 		"aaa",
// 		"aa",
// 		"aaa",
// 	}
// 	utils.Info(">>>", longestCommonPrefix(strs))
// }

// 查找字符串数组中共有的前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i, v := range strs {
		if i == 0 {
			continue
		}
		end := len(prefix)
		l := len(v)
		if l == 0 || end == 0 {
			return ""
		}
		if end > l {
			end = l
			prefix = prefix[:end]
		}
		for ii := 0; ii < end; ii++ {
			if v[ii] != prefix[ii] {
				if ii == 0 {
					return ""
				} else {
					prefix = v[:ii]
					break
				}
			}
		}
	}
	return prefix
}

// 罗马数字转整数 1~3999
// I - 1
// V - 5
// X - 10
// L - 50
// C - 100
// D - 500
// M - 1000
func romanToInt(s string) int {
	l := len(s)
	mm := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	if l < 2 {
		return mm[s]
	}
	sum := 0
	for i := 0; i < l; i++ {
		if i == l-1 {
			sum += mm[string(s[i])]
			break
		}
		n1 := mm[string(s[i])]
		n2 := mm[string(s[i+1])]
		if n1 < n2 {
			sum = sum - n1 + n2
			i++
		} else {
			sum += n1
		}
	}
	return sum
}

// 判断数字是否为回文
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var y int
	for {
		y = y*10 + x%10
		if y == 0 {
			return false
		}
		if x == y {
			return true
		}
		x = x / 10
		if x == y {
			return true
		}
		if y > x {
			return false
		}
	}
	return false
}

// 翻转数字
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	// 是否为负数
	flag := false
	if x < 0 {
		flag = true
		x = 0 - x
	}
	s := ""
	for {
		s += strconv.Itoa(x % 10)
		x = x / 10
		if x == 0 {
			res, _ := strconv.Atoi(s)
			if flag {
				res = 0 - res
			}
			if res > math.MaxInt32 || res < math.MinInt32 {
				return 0
			}
			return res
		}
	}
	return 0
}

// P   A   H   N
// A P L S I I G
// Y   I   R

// 1     5     9     13
// 2  4  6  8  10 12 14
// 3     7     11

// 1     2     3     4
// 5  6  7  8  9  10 11
// 12    13    14

// //q6: convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
// func main() {
// 	utils.Info(">>>>", convert("PAYPALISHIRING", 3))
// }

// func convert(s string, numRows int) string {
// 	l := len(s)
// 	if l < 2 {
// 		return s
// 	}
// 	num := 2*numRows - 2
// 	// 计算有几列
// 	x := l / num //完整个数
// 	y := l % num
// 	head := x
// 	foot := x
// 	if y != 0 {
// 		head++
// 		if y >= numRows {
// 			foot++
// 		}
// 	}
// 	b := make([]byte, l)
// 	for i, v := range s {
// 		index := -1
// 		// 确定位置
// 		m := i % num
// 		n := i / num
// 		if m != 0 {
// 			n++
// 		}
// 		if m < numRows { //竖列
// 			if m == 1 { //第一行
// 				index = n
// 			} else if m == numRows { //最后一行
// 				if y >= numRows {
// 					index = l - 1 - (x - n)
// 				}
// 			} else { //中间几行

// 			}
// 		} else { //斜

// 		}
// 		b[index] = v
// 	}
// 	return string(b)
// }

// q5: 求字符串中最长的回文
// e1:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// e2:
// Input: "cbbd"
// Output: "bb"
// func main() {
// 	utils.Info(">>>>", longestPalindrome("abacded"))
// }

//求字符串中最长的回文
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	maxLen := 1
	for i := 0; i < l-(maxLen/2+maxLen%2); i++ {

	}
	return ""
}

// func longestPalindrome(s string) string {
// 	slen := len(s)
// 	if slen < 2 {
// 		return s
// 	}
// 	utils.Info("===slen===", slen)
// 	maxLen, maxLeft := 1, 0
// 	for i := 0; (i < slen) && (slen-i > (maxLen / 2)); { //为啥要除2
// 		utils.Info("===i===", i)
// 		utils.Info("===maxLen===", maxLen)
// 		left, right := i, i
// 		for right < slen-1 && s[right] == s[right+1] {
// 			right++
// 		}
// 		utils.Info("===right===", right)
// 		i = right + 1

// 		for right < slen-1 && left > 0 && s[right+1] == s[left-1] {
// 			right++
// 			left--
// 		}
// 		utils.Info("===right---", right)
// 		utils.Info("===left---", left)
// 		if right-left+1 > maxLen {
// 			maxLen = right - left + 1
// 			maxLeft = left
// 		}
// 	}
// 	return s[maxLeft : maxLeft+maxLen]
// }

// //求字符串中最长的回文 todo 1
// func longestPalindrome(s string) string {
// 	l := len(s)
// 	if l < 2 {
// 		return s
// 	}
// 	b1 := []byte(s)
// 	b2 := reverse2(b1)
// 	longest := b1[0:1]
// 	longestLen := 0
// 	for i, v := range b1 { // 遍历正序byte数组
// 		if i == l-1 {
// 			break
// 		}
// 		ss := s[i+1:]
// 		if longestLen >= l-i { //剩余长度不足
// 			break
// 		}
// 		n := 0
// 		for n < l-i {
// 			n++
// 			utils.Info(ss)
// 			utils.Info(string(v))
// 			index := strings.LastIndex(ss, string(v))
// 			utils.Info("==index==", index)
// 			if index != -1 {
// 				index = index + i + 1
// 				if index-i+1 > longestLen {
// 					if string(b1[i:index+1]) == string(b2[l-1-index:l-i]) {
// 						longest = b1[i : index+1]
// 						longestLen = index - i + 1
// 						break
// 					}
// 				}
// 				ss = string(b1[i+1 : index])
// 				continue
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	utils.Info(string(longest))
// 	return string(longest)
// }

// 倒序
// func reverse2(b1 []byte) []byte {
// 	l := len(b1)
// 	b2 := make([]byte, l)
// 	for i := 0; i < l; i++ {
// 		b2[i] = b1[l-1-i]
// 	}
// 	return b2
// }

// func main() {
// 	reverse1([]byte("abcde"))
// }

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}

func reverse1(b1 []byte) []byte {
	l := len(b1)
	b2 := make([]byte, l)
	b2 = b1
	utils.Info(b1)
	for i := 0; i < l/2; i++ {
		b2[i], b2[l-1-i] = b1[l-1-i], b1[i]
	}
	utils.Info(b1)
	utils.Info(&b1 == &b2)

	slice1 := (*Slice)(unsafe.Pointer(&b1))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice1.ptr, slice1.len, slice1.cap)
	slice2 := (*Slice)(unsafe.Pointer(&b2))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice2.ptr, slice2.len, slice2.cap)
	return b2
}

// q4:
// 求两个数组的中位数 todo
// e1:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5
// e2:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0
// O(log (m+n))
// func main() {
// 	nums1 := []int{8, 1, 3}
// 	nums2 := []int{2}
// 	utils.Info(findMedianSortedArrays(nums1, nums2)) //4.5
// }

// 基本需求
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var res float64
	if len(nums2) != 0 {
		nums1 = append(nums1, nums2...)
	}
	l := len(nums1)
	if l == 0 {
		return 0
	}
	sort.Ints(nums1)
	i := l / 2
	if l%2 == 0 {
		res = float64(nums1[i-1]+nums1[i]) / 2
	} else {
		res = float64(nums1[i])
	}
	return res
}

//q3:
// "abcabcbb"-->"abc"
// "bbbbb"-->"b"
// "pwwkew"-->"wke"
// func main() {
// 	str := "pwwkewew"
// 	utils.Info(">>>>>", lengthOfLongestSubstring(str))
// }

//未测试 todo
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}
	var max, max1, start int
	lastMap := make(map[rune]int, 0)
	for i, v := range s {
		if m, ok := lastMap[v]; ok {
			// 存在该字母
			// 判断下标是否大于开始下标
			if m > start {
				utils.Info(i)
				max1 = i - m
				start = m + 1
			} else if m == start {
				utils.Info(i)
				max1 = i - m
				start = m + 1
			} else {
				utils.Info(i)
				max1 = i - start + 1
			}
		} else {
			utils.Info(i)
			// 不存在该字母
			max1 = i - start + 1
		}
		if max1 > max {
			utils.Info(i)
			max = max1
		}
		// 修改字母最后出现下标
		lastMap[v] = i
	}
	return max
}

// func lengthOfLongestSubstring(s string) int {
// 	character := [128]int{}
// 	max, count, last := 0, 0, 0
// 	for i, v := range s {
// 		if character[v] < last {
// 			count += 1
// 		} else {
// 			last = character[v]
// 			if count > max {
// 				max = count
// 			}
// 			count = i + 1 - last
// 		}
// 		character[v] = i + 1
// 	}
// 	if count > max {
// 		max = count
// 	}
// 	return max
// }

// func lengthOfLongestSubstring(s string) int {
// 	l := len(s)
// 	if l < 2 {
// 		return l
// 	}
// 	last := 0
// 	max := 0
// 	m := map[byte]int{}
// 	for i := 0; i < l; i++ {
// 		utils.Info("iiiii", i)
// 		n := 0
// 		if v, ok := m[s[i]]; ok {
// 			utils.Info("vvvvv  ", v)
// 			last = v
// 			n = i - last
// 			if s[i] == s[last] {
// 				n = i - last
// 			} else {
// 				n = i - last + 1
// 			}
// 		} else {
// 			n = i - last + 1
// 			utils.Info("====")
// 		}
// 		utils.Info(last)
// 		if n > max {
// 			max = n
// 		}
// 		utils.Info("mmmm  ", max)
// 		m[s[i]] = i
// 	}
// 	return max
// }

// func lengthOfLongestSubstring(s string) int {
// 	l := len(s)
// 	if l < 2 {
// 		return l
// 	}
// 	max := 0
// 	for i := 0; i < l-max; i++ {
// 		f := false
// 		m := map[byte]int{s[i]: i}
// 		for j := i + 1; j < l; j++ {
// 			if _, ok := m[s[j]]; ok {
// 				n := j - i
// 				if n > max {
// 					max = n
// 					utils.Info(max)
// 				}
// 				f = true
// 				break
// 			} else {
// 				m[s[j]] = j
// 			}
// 		}
// 		if !f {
// 			max = l - i
// 		}
// 	}
// 	return max
// }

// func lengthOfLongestSubstring(s string) int {
// 	l := len(s)
// 	if l <= 1 {
// 		return l
// 	}
// 	max := 1
// 	for i := 0; i < l-1; i++ {
// 		m := map[byte]int{s[i]: i}
// 		for j := i + 1; j < l; j++ {
// 			if v, ok := m[s[j]]; ok {
// 				n := j - v
// 				if n > max {
// 					max = n
// 				}
// 				break
// 			} else {
// 				m[s[j]] = j
// 			}
// 		}
// 	}
// 	return max
// }

//q2:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807
// func main() {
// 	l1 := &ListNode{
// 		Val: 9,
// 		Next: &ListNode{
// 			Val: 9,
// 			Next: &ListNode{
// 				Val:  9,
// 				Next: nil,
// 			},
// 		},
// 	}
// 	l2 := &ListNode{
// 		Val: 1,
// 		// Next: &ListNode{
// 		// 	Val: -6,
// 		// Next: &ListNode{
// 		// 	Val:  4,
// 		Next: nil,
// 		// },
// 		// },
// 	}
// 	res := addTwoNumbers(l1, l2)
// 	for {
// 		utils.Info(res.Val)
// 		if res.Next == nil {
// 			return
// 		}
// 		res = res.Next
// 	}
// 	return
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{
		Val:  l1.Val + l2.Val,
		Next: addTwoNumbers(l1.Next, l2.Next),
	}
	return CheckNum(res)
}

//q2:改
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	carry := 0
// 	res := &ListNode{}
// 	r := res
// 	for l1 != nil || l2 != nil {
// 		sum := carry
// 		if l1 != nil {
// 			sum += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			sum += l2.Val
// 			l2 = l2.Next
// 		}
// 		carry = sum / 10
// 		r.Next = &ListNode{}
// 		r = r.Next
// 		r.Val = sum % 10
// 	}
// 	if carry != 0 {
// 		r.Next = &ListNode{}
// 		r.Next.Val = carry
// 	}
// 	return res.Next
// }

func CheckNum(res *ListNode) *ListNode {
	a := res.Val / 10
	if a != 0 {
		b := res.Val % 10
		if res.Next == nil {
			res = &ListNode{
				Val: b,
				Next: &ListNode{
					Val:  a,
					Next: nil,
				},
			}
		} else {
			res = &ListNode{
				Val: b,
				Next: &ListNode{
					Val:  res.Next.Val + a,
					Next: res.Next.Next,
				},
			}
		}
		if res.Next != nil && res.Next.Val/10 != 0 {
			res.Next = CheckNum(res.Next)
		}
	}
	return res
}

//q1:

// func main() {
// 	nums := []int{-1, -2, -3, -4, -5}
// 	res := twoSum(nums, -8)
// 	utils.Info(res)
// }

func twoSum(nums []int, target int) []int {
	res := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return res
}

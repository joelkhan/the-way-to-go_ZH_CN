package main

import "fmt"

func main() {
	// ex7.12 注意，这一章有两个练习7.12
	str := "HIJKLMNabcdefg"
	s1, s2 := ex712(str, 2)
	fmt.Printf("sub str1: %s\n", s1)
	fmt.Printf("sub str2: %s\n", s2)

	// ex7.13
	ex713()

	// ex7.14
	ex714()

	// ex7.15
	ex715()

	// ex7.16
	ex716()

	// ex7.17
	ex717()
}

// 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func ex712(str string, i int) (s1, s2 string) {
	if i >= 0 && i <= len(str) {
		s1 = str[:i]
		s2 = str[i:]
	}
	return
}

// 假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
func ex713() {
	str := "HIJKLMNabcdefg"
	fmt.Printf("ex713, origin: %s\n", str)
	fmt.Printf("ex713, formatted: %s\n", str[len(str)/2:]+str[:len(str)/2])
}

/* 1. 使用 []byte 类型的切片反转字符串，即将 "Google" 转换成 "elgooG"
 *    如果您使用两个切片来实现反转，请再尝试使用一个切片
 * 2. 如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。
 */
func ex714() {
	str := "Google"
	sli := []byte(str)
	fmt.Printf("string %s, len: %d, cap: %d\n", sli, len(sli), cap(sli))
	for i := 0; i < len(sli)/2; i++ {
		//fmt.Printf("%c ", sli[i])
		sli[i], sli[len(sli)-1-i] = sli[len(sli)-1-i], sli[i]
	}
	fmt.Printf("reverse: %s\n", sli)

	str2 := "jsこん dk"
	sli2 := []int32(str2)
	fmt.Printf("%-10s |%10s| len: %d, cap: %d\n", "origin:", str2, len(sli2), cap(sli2))
	for i := 0; i < len(sli2)/2; i++ {
		sli2[i], sli2[len(sli2)-1-i] = sli2[len(sli2)-1-i], sli2[i]
	}
	fmt.Printf("%-10s |%10s|\n", "reverse:", string(sli2))
}

// 遍历一个字符数组，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
func ex715() {
	//str := "let's see what will happend about golang in the future!"
	str := "asddffff About  Golang~"
	sli := []byte(str)
	var res []byte // == nil
	var lastOne byte = 0
	for _, c := range sli {
		if c != lastOne {
			//fmt.Printf("%c ", c)
			res = append(res, c)
		}
		lastOne = c

	}
	fmt.Printf("origin: %s\n", sli)
	fmt.Printf("unique: %s\n", res)
}

// 使用冒泡排序的方法排序一个包含整数的切片（算法的定义可参考 维基百科）。
func ex716() {

}

/* 编写一个函数 mapFunc 要求接受以下 2 个参数：
 *  - 一个将整数乘以 10 的函数
 *  - 一个整数列表
 * 最后返回保存运行结果的整数列表。
 */
func ex717() {

}

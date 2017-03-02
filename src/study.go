package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("hello go word")

	// num++为语句不是表达式;
	// 不支持 x = num++;
	// 不支持 ++num
	//var num int
	//num++
	//fmt.Println(num)

	//// for 默认实现
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = "\n"
	//}
	//fmt.Println(s)

	// for 只有条件
	//var condition int
	//for condition < 10 {
	//	condition++
	//	fmt.Println(condition)
	//}

	//// 无限循环
	//for {
	//	fmt.Println("~~~~~~~~~~~~~~~~~")
	//}

	// 省略符号_
	//s, sep := "",""
	//
	//for _, arg := range os.Args[1:]{
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)

	//fmt.Println(strings.Join(os.Args[1:], " "))
	//fmt.Println(os.Args[:])

	// go执行mac 命令
	//exec.Command("open" ,"/Users/hangsong/Desktop/书").Start()

	// dup1
	// 输入，使用 command + D 结束输入
	// map[key]value
	//counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//
	//for input.Scan() {
	//	counts[input.Text()]++
	//	fmt.Println("__________")
	//}
	//
	//for line, n := range counts {
	//	if (n > 1) {
	//		fmt.Println("%d\t%s\n", n, line)
	//	}
	//}

	// dup2
	//counts := make(map[string]int)
	//files := os.Args[1:]
	//// 输入参数为空
	//if len(files) == 0 {
	//	countLines(os.Stdin, counts)
	//} else {
	//	for _, arg := range files {
	//		f, err := os.Open(arg)
	//		// 打开错误
	//		if err != nil {
	//			fmt.Fprint(os.Stderr, err)
	//			continue
	//		}
	//		countLines(f, counts)
	//		f.Close()
	//	}
	//}
	//
	//for line, n := range counts {
	//	if (n > 1) {
	//		fmt.Println("%d\t%s\n", n, line)
	//	}
	//}

	// dup3 读入文件统计当前行出现次数
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}

}

//func countLines(f *os.File, counts map[string]int) {
//	input := bufio.NewScanner(f);
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//}

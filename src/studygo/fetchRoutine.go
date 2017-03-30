package main

import (
	"time"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
)

/**
并发获取多个URL
 */
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs 耗时\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now();
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint("错误： ", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprint("%.2fs %7d %s", secs, nbytes, url)
}

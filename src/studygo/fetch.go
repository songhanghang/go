package studygo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "featch: %v '第一'\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch:  reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Print("%s", b)
	}
}

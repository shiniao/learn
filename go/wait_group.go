// wait group 用来等待一组goroutines完成

package learn

import (
	"fmt"
	"sync"
)

type HTTPPkg struct{}

func (HTTPPkg) Get(url string) string {
	return fmt.Sprintf("get %s ok.", url)
}

var http HTTPPkg

func wg() {

	var wg sync.WaitGroup
	
	urls := []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		// 每次加入一个goroutine
		wg.Add(1)

		go func(url string) {
			// 等待完成
			defer wg.Done()
			// 执行程序
			result := http.Get(url)
			fmt.Println(result)
		}(url)
	}

	wg.Wait()

}

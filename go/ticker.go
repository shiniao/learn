package learn

import (
	"fmt"
	"time"
)

// Ticker 代表一个ticker
func Ticker() {
	// 设置一个一秒的ticker
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// 接收完成信号
	done := make(chan bool)

	go func() {
		// 十秒之后结束
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("current time: ", t)
		case <-done:
			fmt.Println("Done!")
			return
		}
	}
}

package learn

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

// goroutine 不能无限创建，过度会大量占用系统资源，导致被系统 kill
func TestGoroutineMore(t *testing.T) {
	taskMany := math.MaxInt64

	for i := 0; i < taskMany; i++ {
		go func(i int) {
			fmt.Println("go func", i, "goroutine count: ", runtime.NumGoroutine())
		}(i)
	}
	// 	panic: too many concurrent operations on a single file or socket (max 1048575)
}

// 解决办法：利用无缓冲channel与任务发送/执行分离方式
var group = sync.WaitGroup{}

// 用户业务
func busi(ch chan int) {
	for c := range ch {
		fmt.Println("go task: ", c, "goroutine count:", runtime.NumGoroutine())
		group.Done()
	}
}

// send task to channel
func sendTask(task int, ch chan int) {
	group.Add(1)
	ch <- task
}
func TestGoroutineMoreSolution(t *testing.T) {
	ch := make(chan int)
	// 启动 goroutine
	for i := 0; i < 3; i++ {
		go busi(ch)
	}

	// send task
	for i := 0; i < 10000; i++ {
		sendTask(i, ch)
	}

	group.Wait()
}

// 最简单的方式，使用sync 同步机制, waitGroup
var wg2 = sync.WaitGroup{}

func busi2(task int) {
	fmt.Println("go task: ", task, "goroutine count:", runtime.NumGoroutine())
	wg2.Done()

}

/*
task  ---\                       /-->   goroutine1
task  --->  [task|task|task|...] --->   goroutine2
task  ---/        channel        \-->   goroutine3
...

*/
func TestGoroutineMoreUseChannel(t *testing.T) {
	for i := 0; i < 10000; i++ {
		wg2.Add(1)
		go busi2(i)
	}
	wg2.Wait()
}

// signal 优雅的关闭信号
// 通过channel发送信号，然后注册接收信号到channel

package learn

import (
	"fmt"
	"os"
	"os/signal"
)

func mySignal() {
	terminate := make(chan os.Signal, 1)
	// 注册interrupt信号到 terminate
	signal.Notify(terminate, os.Interrupt)

	// 接收信号
	term := <-terminate

	fmt.Println("\nSignal: ", term)

}

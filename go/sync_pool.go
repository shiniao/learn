// sync.pool
// 当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，
// 导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。
// 形成 “并发⼤－占⽤内存⼤－GC 缓慢－处理并发能⼒降低－并发更⼤”这样的恶性循环。
// pool会复用对象，减少 GC 压力。

package learn

import (
	"sync"
)

// pool 中返回了 httpPkg 结构体
// 可以复用
var pool = sync.Pool{
	New: func() interface{} {
		return new(HTTPPkg)
	},
}


// Pool 减少了GC的压力，pool中的对象可以复用，减少内存压力
func syncPool(url string) string {
	// 从pool取一个
	http := pool.Get().(*HTTPPkg)
	result := http.Get(url)
	// 用了之后放回去
	pool.Put(http)
	
	return result

}

// 不使用 pool
func syncNoPool(url string) string {
	http := HTTPPkg{}
	return http.Get(url)
}

// sync.pool
// 当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，
// 导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。
// 形成 “并发⼤－占⽤内存⼤－GC 缓慢－处理并发能⼒降低－并发更⼤”这样的恶性循环。
// pool会复用对象，减少 GC 压力。

package learn

import (
	"sync"
)

//Small just small
type Small struct {
	a int
}

// pool 中返回了 Small 结构体, 可以复用
var pool = sync.Pool{
	New: func() interface{} {
		return new(Small)
	},
}

// NewSmallPool init a small pool
func NewSmallPool() *Small {
	return pool.Get().(*Small)
}

//Free 释放对象
func (s *Small) Free() {
	pool.Put(s)
}

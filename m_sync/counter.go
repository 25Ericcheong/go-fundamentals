package sync

import "sync"

// embedding mutex in struct is wrong and bad
//type Counter struct {
//	sync.Mutex
//	value int
//}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// embedding sync.Mutex into struct will look like this
// type Counter struct {
// 	sync.Mutex
// 	value int
// }

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

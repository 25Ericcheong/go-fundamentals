package main

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

// this means that when each go routine (in our case; a 1000) calls this method - it must acquire a lock first and all other go routines will have to wait for it to be unlocked before getting access
func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
func (c *Counter) Value() int {
	return c.value
}

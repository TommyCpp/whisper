package model

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
	instance *interface{}
	once     Once
}

type Once struct {
	mu   sync.Mutex
	done int32
}

func (o *Once) Do(f func()) {
	//implement a check-lock-check mechanism
	if atomic.LoadInt32(&o.done) == 1 {
		return
	} else {
		o.mu.Lock()
		defer o.mu.Unlock()
		if o.done == 0 {
			defer atomic.StoreInt32(&o.done, 1)
			f()
		}
	}
}

func (s *Singleton) getInstance() interface{} {
	s.once.Do(func() {
		s.Init()
	})
	return s.instance
}
func (s *Singleton) Init() *interface{} {
	return nil
}

package di

import (
	"sync"
)

var Di = &di{
	store: make(map[string]any),
	mutex: sync.RWMutex{},
}

type di struct {
	store map[string]any
	mutex sync.RWMutex
}

func (d *di) SetDi(name string, v any) {
	d.mutex.Lock()
	d.store[name] = v
	d.mutex.Unlock()
}

func (d *di) GetDi(name string) any {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	v := d.store[name]
	return v
}

func Get(name string) any {
	return Di.GetDi(name)
}

func Set(name string, v any) {
	Di.SetDi(name, v)
}

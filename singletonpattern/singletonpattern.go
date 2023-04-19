package singletonpattern

import (
	"fmt"
	"sync"
)

type singleton struct {
}

func (*singleton) show() {
	fmt.Println("This is a singleton")
}

var instance_eager *singleton = &singleton{}
var instance_lazy *singleton

func getInstanceEager() *singleton {
	return instance_eager
}

var once sync.Once

func getInstanceLazy() *singleton {
	once.Do(func() {
		if instance_lazy == nil {
			instance_lazy = &singleton{}
		}
	})
	return instance_lazy
}

package singletonpattern

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var instanceObj *singleton

	instanceObj = new(singleton)

	fmt.Println("Eager style...")
	instanceObj = getInstanceEager()
	instanceObj.show()

	fmt.Println("Lazy style...")
	instanceObj = getInstanceLazy()
	instanceObj.show()
}

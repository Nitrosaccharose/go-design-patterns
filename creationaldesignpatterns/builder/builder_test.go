package builder

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	director1 := Director{&Builder1{}}
	house1 := director1.buildHouse()
	fmt.Println(house1)
}

func Test2(t *testing.T) {
	director2 := Director{&Builder2{}}
	house2 := director2.buildHouse()
	fmt.Println(house2)
}

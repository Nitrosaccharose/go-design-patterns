package prototype

import (
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	computer1 := &computer{serialNumber: strconv.Itoa(1234567890)}
	computer2 := computer1.clone()
	fmt.Println(computer1)
	fmt.Println(computer2)

}

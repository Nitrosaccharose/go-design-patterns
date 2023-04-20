package adapter

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	stringReverser := &StringReverser{}
	fmt.Println(stringReverser.reverse("hello"))

	numberAdapter := &NumberAdapter{reverser: stringReverser}
	fmt.Println(numberAdapter.reverse(1234567890))
}

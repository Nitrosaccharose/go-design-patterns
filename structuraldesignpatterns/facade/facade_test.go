package facade

import "testing"

func Test1(t *testing.T) {
	hf := HomeFacade{}
	hf.OfficeMode()
}

func Test2(t *testing.T) {
	hf := HomeFacade{}
	hf.RestMode()
}

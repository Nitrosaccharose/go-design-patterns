package facade

import "fmt"

// Light 灯
type Light struct {
}

func (*Light) on() {
	fmt.Println("Light is on")
}
func (*Light) off() {
	fmt.Println("Light is off")
}

// TV 电视
type TV struct {
}

func (*TV) on() {
	fmt.Println("TV is on")
}
func (*TV) off() {
	fmt.Println("TV is off")
}

type Computer struct {
}

func (*Computer) on() {
	fmt.Println("Computer is on")
}

func (*Computer) off() {
	fmt.Println("Computer is off")
}

// AirConditioner 空调
type AirConditioner struct {
}

func (*AirConditioner) on() {
	fmt.Println("AirConditioner is on")
}

func (*AirConditioner) off() {
	fmt.Println("AirConditioner is off")
}

// HomeFacade 家庭模式控制器
type HomeFacade struct {
	light1         Light
	light2         Light
	tv             TV
	computer       Computer
	airConditioner AirConditioner
}

func (hf *HomeFacade) OfficeMode() {
	fmt.Println("--- Office Mode ---")
	hf.light1.on()
	hf.light2.on()
	hf.computer.on()
}

func (hf *HomeFacade) RestMode() {
	fmt.Println("--- Rest Mode ---")
	hf.light2.on()
	hf.tv.on()
	hf.airConditioner.on()
}

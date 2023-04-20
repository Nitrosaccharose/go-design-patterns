package abstractpattern

import "testing"

func Test1(t *testing.T) {
	var factory AbstractFactory
	factory = new(ChinaFactory)
	chinaApple := factory.CreateApple()
	chinaBanana := factory.CreateBanana()
	chinaPair := factory.CreatePear()
	chinaApple.show()
	chinaBanana.show()
	chinaPair.show()

	factory = new(USAFactory)
	usaApple := factory.CreateApple()
	usaBanana := factory.CreateBanana()
	usaPair := factory.CreatePear()
	usaApple.show()
	usaBanana.show()
	usaPair.show()

	factory = new(JapanFactory)
	japanApple := factory.CreateApple()
	japanBanana := factory.CreateBanana()
	japanPair := factory.CreatePear()
	japanApple.show()
	japanBanana.show()
	japanPair.show()
}

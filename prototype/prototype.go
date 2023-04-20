package prototype

type Cloneable interface {
	clone() Cloneable
}

type computer struct {
	serialNumber string
	Cloneable
}

func (c *computer) clone() Cloneable {
	return &computer{serialNumber: c.serialNumber}
}

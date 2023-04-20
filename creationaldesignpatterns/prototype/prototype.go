package prototype

type CloneAble interface {
	clone() CloneAble
}

type computer struct {
	serialNumber string
	CloneAble
}

func (c *computer) clone() CloneAble {
	return &computer{serialNumber: c.serialNumber}
}

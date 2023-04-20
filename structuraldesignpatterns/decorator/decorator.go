package decorator

// Food 汉堡
type Food interface {
	getPrice() int
}

// Hamburger 汉堡
type Hamburger struct {
	Food
}

func (h *Hamburger) getPrice() int {
	return 2
}

// Sandwich 三明治
type Sandwich struct {
	Food
}

func (s *Sandwich) getPrice() int {
	return 1
}

type JardiniereDecorator struct {
	food Food
}

// Meat 肉饼
type Meat struct {
	JardiniereDecorator
}

func (m *Meat) getPrice() int {
	return m.food.getPrice() + 10
}

// Cheese 奶酪
type Cheese struct {
	JardiniereDecorator
}

func (c *Cheese) getPrice() int {
	return c.food.getPrice() + 5
}

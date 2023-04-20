package builder

// Builder 生成器
type Builder interface {
	buildWall()
	buildRoof()
	buildFloor()
	getHouse() House
}

// House 房子
type House struct {
	wall  string
	roof  string
	floor string
	Builder
}

// Builder1 具体生成器1
type Builder1 struct {
	house House
}

func (b *Builder1) buildWall() {
	b.house.wall = "red"
}

func (b *Builder1) buildRoof() {
	b.house.roof = "blue"
}

func (b *Builder1) buildFloor() {
	b.house.floor = "yellow"
}

func (b *Builder1) getHouse() House {
	return b.house
}

// Builder2 具体生成器2
type Builder2 struct {
	house House
}

func (b *Builder2) buildWall() {
	b.house.wall = "white"
}

func (b *Builder2) buildRoof() {
	b.house.roof = "black"
}

func (b *Builder2) buildFloor() {
	b.house.floor = "green"
}

func (b *Builder2) getHouse() House {
	return b.house
}

// Director 指挥者
type Director struct {
	builder Builder
}

func (d *Director) Director(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.buildWall()
	d.builder.buildRoof()
	d.builder.buildFloor()
	return d.builder.getHouse()
}

package wires

type coordinate string

func newCoordinate(...int) string {
	
}

type cell struct {
	value int
	coordinate
	up, left, down, right *cell
}

type grid struct {
	head     *cell
	segments map[coordinate]*cell
}

func newGrid() *grid {
	return &grid{
		head: &cell{
			value:      1,
			coordinate: "0,0",
		},
	}
}

func (g *grid) Up() {
	if exist, ok := g.segments[g.head.coordinate]; ok {
		// modify an existing cell
		exist.value++
		g.head = exist
	} else {
		// create a new cell
		g.head = g.head.move(0, 1)
	}

}

func (g *grid) left() {

}

func (c *cell) move(x, y int) *cell {
	newCell := &cell{
		value:      c.value,
		coordinate: "",
		up:         c.up,
		left:       c.left,
		down:       c.down,
		right:      c.right,
	}
	return newCell
}

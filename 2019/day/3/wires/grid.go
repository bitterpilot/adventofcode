package wires

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type coordinate struct {
	X, Y int
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c *coordinate) Add(nc coordinate) {
	c.X = c.X + nc.X
	c.Y = c.Y + nc.Y
}

type wire struct {
	ID int
}

type grid struct {
	head   coordinate
	wireID int // the current wire
	world  map[coordinate]wire

	intersections []coordinate
}

// directions
const (
	U int = 1
	D int = -1
	R int = 1
	L int = -1
)

func Solve(wire1, wire2 []string) int {
	grid := NewGrid()
	grid.DrawWire(wire1)
	grid.DrawWire(wire2)
	return grid.FindDistance()
}

func NewGrid() grid {
	return grid{
		head: coordinate{
			X: 0,
			Y: 0,
		},
		world: map[coordinate]wire{
			{
				X: 0,
				Y: 0,
			}: {},
		},
	}
}

// FindDistance reports the shortest distance between the origin(0,0) and the
// closest intersection.
func (g *grid) FindDistance() int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	distances := make([]int, 0, cap(g.intersections))
	for _, v := range g.intersections {
		distances = append(distances, abs(v.X)+abs(v.Y))
	}
	sort.Ints(distances)
	return distances[0]
}

// DrawWire handles multi direction movements
// it takes a slice of movements ie ["R2","D2"]
func (g *grid) DrawWire(wire []string) error {
	g.ResetHead()
	g.wireID = rand.Int()
	for _, move := range wire {
		err := g.Move(move)
		if err != nil {
			return err
		}
	}

	return nil
}

// ResetHead sets the head postingion of the wire back to 0,0
func (g *grid) ResetHead() {
	g.head = coordinate{
		X: 0,
		Y: 0,
	}
}

// Move handles single direction movements
func (g *grid) Move(movement string) error {
	move, err := movementParser(movement)
	if err != nil {
		return fmt.Errorf("movement parser: %w", err)
	}

	for _, step := range move {
		g.head.Add(step)

		if w, ok := g.world[g.head]; ok {
			// if it exists check if it is a different wire
			if g.wireID != w.ID {
				g.intersections = append(g.intersections, g.head)
				continue
			}
		}
		// otherwise add it to the map
		g.world[g.head] = wire{g.wireID}

	}
	return nil
}

func movementParser(movement string) ([]coordinate, error) {
	// decode movement in to component parts
	direction := string(movement[0])
	magnitude, err := strconv.Atoi(movement[1:])
	if err != nil {
		return nil, fmt.Errorf("error parsing int: %w", err)
	}

	var co coordinate
	switch direction {
	// X axis
	case "R":
		co = coordinate{X: 1, Y: 0}
	case "L":
		co = coordinate{X: -1, Y: 0}
	// Y axis
	case "U":
		co = coordinate{X: 0, Y: 1}
	case "D":
		co = coordinate{X: 0, Y: -1}
	default:
		return nil, fmt.Errorf(
			"movement parser: error parsing direction, received: %v", direction)
	}

	ret := make([]coordinate, 0, magnitude)

	for i := 0; i < magnitude; i++ {
		ret = append(ret, co)
	}

	return ret, nil
}

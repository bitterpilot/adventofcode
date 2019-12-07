package wires

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_grid_Move(t *testing.T) {
	// BUG: this test was written before the wireID was added
	type fields struct {
		head          coordinate
		world         map[coordinate]wire
		intersections []coordinate
	}
	type args struct {
		movement string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantGrid *fields
	}{
		{
			name: "Simple move",
			fields: fields{
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
				intersections: nil,
			},
			args: args{
				movement: "R1",
			},
			wantErr: false,
			wantGrid: &fields{
				head: coordinate{
					X: 1,
					Y: 0,
				},
				world: map[coordinate]wire{
					{
						X: 0,
						Y: 0,
					}: {},
					{
						X: 1,
						Y: 0,
					}: {},
				},
			},
		}, {
			name: "Simple Chain",
			fields: fields{
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
				intersections: nil,
			},
			args: args{
				movement: "U5",
			},
			wantErr: false,
			wantGrid: &fields{
				head: coordinate{
					X: 0,
					Y: 5,
				},
				world: map[coordinate]wire{
					{
						X: 0,
						Y: 0,
					}: {},
					{
						X: 0,
						Y: 1,
					}: {},
					{
						X: 0,
						Y: 2,
					}: {}, {
						X: 0,
						Y: 3,
					}: {},
					{
						X: 0,
						Y: 4,
					}: {},
					{
						X: 0,
						Y: 5,
					}: {},
				},
			},
		}, {
			name: "Simple Chain Negative",
			fields: fields{
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
				intersections: nil,
			},
			args: args{
				movement: "D5",
			},
			wantErr: false,
			wantGrid: &fields{
				head: coordinate{
					X: 0,
					Y: -5,
				},
				world: map[coordinate]wire{
					{
						X: 0,
						Y: 0,
					}: {},
					{
						X: 0,
						Y: -1,
					}: {},
					{
						X: 0,
						Y: -2,
					}: {}, {
						X: 0,
						Y: -3,
					}: {},
					{
						X: 0,
						Y: -4,
					}: {},
					{
						X: 0,
						Y: -5,
					}: {},
				},
			},
		}, {
			name: "Simple Intersect",
			fields: fields{
				head: coordinate{
					X: 0,
					Y: 0,
				},
				world: map[coordinate]wire{
					{
						X: 0,
						Y: 0,
					}: {}, {
						X: 0,
						Y: 4,
					}: {},
				},
				intersections: nil,
			},
			args: args{
				movement: "U5",
			},
			wantErr: false,
			wantGrid: &fields{
				head: coordinate{
					X: 0,
					Y: 5,
				},
				world: map[coordinate]wire{
					{
						X: 0,
						Y: 0,
					}: {},
					{
						X: 0,
						Y: 1,
					}: {},
					{
						X: 0,
						Y: 2,
					}: {}, {
						X: 0,
						Y: 3,
					}: {},
					{
						X: 0,
						Y: 4,
					}: {},
					{
						X: 0,
						Y: 5,
					}: {},
				},
				intersections: []coordinate{
					{
						X: 0,
						Y: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &grid{
				head:          tt.fields.head,
				world:         tt.fields.world,
				intersections: tt.fields.intersections,
			}
			if err := g.Move(tt.args.movement); (err != nil) != tt.wantErr {
				t.Errorf("grid.Move() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO: Deal with the wireID bug. (maybe just set it to zero)
			if !reflect.DeepEqual(g, tt.wantGrid) {
				t.Error("\ngrids do not match")
				fmt.Printf("want: %+v\n", tt.wantGrid)
				fmt.Printf(" got: %+v\n", g)
			}
		})
	}
}

func Test_grid_DrawWire(t *testing.T) {
	type fields struct {
		head          coordinate
		world         map[coordinate]wire
		intersections []coordinate
	}
	type args struct {
		wire []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantGrid *fields
	}{
		{
			name: "Simple",
			fields: fields{
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
				intersections: nil,
			},
			args: args{
				wire: []string{"R1", "D2", "L3", "U4"},
			},
			wantErr: false,
			wantGrid: &fields{
				head: coordinate{
					X: -2,
					Y: 2,
				},
				world: map[coordinate]wire{
					{X: -2, Y: -2}: {},
					{X: -2, Y: -1}: {},
					{X: -2, Y: 0}:  {},
					{X: -2, Y: 1}:  {},
					{X: -2, Y: 2}:  {},
					{X: -1, Y: -2}: {},
					{X: 0, Y: -2}:  {},
					{X: 0, Y: 0}:   {},
					{X: 1, Y: -2}:  {},
					{X: 1, Y: -1}:  {},
					{X: 1, Y: 0}:   {},
				},
				intersections: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &grid{
				head:          tt.fields.head,
				world:         tt.fields.world,
				intersections: tt.fields.intersections,
			}
			if err := g.DrawWire(tt.args.wire); (err != nil) != tt.wantErr {
				t.Errorf("grid.DrawWire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Examples(t *testing.T) {
	type example struct {
		name         string
		wire1, wire2 []string
		distance     int
	}
	tests := []example{
		{
			name:     "example 1",
			wire1:    []string{"R8", "U5", "L5", "D3"},
			wire2:    []string{"U7", "R6", "D4", "L4"},
			distance: 6,
		},
		{
			name:     "example 2",
			wire1:    []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			wire2:    []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			distance: 159,
		},
		{
			name:     "example 3",
			wire1:    []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			wire2:    []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			distance: 135,
		},
	}

	for _, test := range tests {
		dist := Solve(test.wire1, test.wire2)

		if dist != test.distance {
			t.Errorf("%s failed: want %d, got %d", test.name, test.distance, dist)
		}
	}
}

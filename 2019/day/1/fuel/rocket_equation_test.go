package fuel_test

import (
	"testing"

	"github.com/bitterpilot/adventofcode/2019/day/1/fuel"
)

func TestFuelCounterUpper(t *testing.T) {
	type args struct {
		masses []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example 1",
			args: args{
				masses: []float64{12},
			},
			want: 2,
		}, {
			name: "example 2",
			args: args{
				masses: []float64{14},
			},
			want: 2,
		}, {
			name: "example 3",
			args: args{
				masses: []float64{1969},
			},
			want: 966,
		}, {
			name: "example 4",
			args: args{
				masses: []float64{100756},
			},
			want: 50346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuel.CounterUpper(tt.args.masses); got != tt.want {
				t.Errorf("FuelCounterUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

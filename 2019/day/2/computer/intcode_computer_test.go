package computer_test

import (
	"reflect"
	"testing"

	"github.com/bitterpilot/adventofcode/2019/day/2/computer"
)

func TestIntcode(t *testing.T) {
	type args struct {
		in0 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 2",
			args: args{
				in0: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
			want: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		}, {
			name: "Example 3",
			args: args{
				in0: []int{1, 0, 0, 0, 99},
			},
			want: []int{2, 0, 0, 0, 99},
		}, {
			name: "Example 4",
			args: args{
				in0: []int{2, 3, 0, 3, 99},
			},
			want: []int{2, 3, 0, 6, 99},
		}, {
			name: "Example 5",
			args: args{
				in0: []int{2, 4, 4, 5, 99, 0},
			},
			want: []int{2, 4, 4, 5, 99, 9801},
		}, {
			name: "Example 6",
			args: args{
				in0: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computer.Intcode(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

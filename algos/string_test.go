package algos

import (
	"reflect"
	"testing"
)

func TestToColumnSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "single line -",
			args: args{
				"1 2 3",
			},
			want: [][]string{
				{"1"},
				{"2"},
				{"3"},
			},
		},
		{
			name: "single line |",
			args: args{
				`1
2
3`,
			},
			want: [][]string{
				{"1", "2", "3"},
			},
		},
		{
			name: "double line",
			args: args{
				`1 2
3 4
5 6`,
			},
			want: [][]string{
				{"1", "3", "5"},
				{"2", "4", "6"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToColumnSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToColumnSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

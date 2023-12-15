package structs

import (
	"reflect"
	"testing"
)

func TestNonogram_MakeAllPossible(t *testing.T) {
	type fields struct {
		Line        []byte
		Pattern     []int
		NumPossible int
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]byte
	}{
		{
			name: "test1",
			fields: fields{
				Line:        []byte("???.###"),
				Pattern:     []int{1, 1, 3},
				NumPossible: 1,
			},
			want: [][]byte{
				[]byte("#.#.###"),
			},
		},
		{
			name: "test2",
			fields: fields{
				Line:        []byte("?#?#?#?#?#?#?#?"),
				Pattern:     []int{1, 3, 1, 6},
				NumPossible: 1,
			},
			want: [][]byte{
				[]byte(".#.###.#.######"),
			},
		},
		{
			name: "test2",
			fields: fields{
				Line:        []byte("????.######..#####."),
				Pattern:     []int{1, 6, 5},
				NumPossible: 1,
			},
			want: [][]byte{
				[]byte("#....######..#####."),
				[]byte(".#...######..#####."),
				[]byte("..#..######..#####."),
				[]byte("...#.######..#####."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Nonogram{
				Line:        tt.fields.Line,
				Pattern:     tt.fields.Pattern,
				NumPossible: tt.fields.NumPossible,
			}
			if got := n.MakeAllPossible(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeAllPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonogram_Unfold(t *testing.T) {
	type fields struct {
		Line        []byte
		Pattern     []int
		NumPossible int
	}
	type args struct {
		c int
	}
	type wants struct {
		line    []byte
		pattern []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   wants
	}{
		{
			name: "test1",
			fields: fields{
				Line:        []byte("???.###"),
				Pattern:     []int{1, 1, 3},
				NumPossible: 0,
			},
			args: args{
				c: 5,
			},
			want: wants{
				line:    []byte("???.###????.###????.###????.###????.###"),
				pattern: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Nonogram{
				Line:        tt.fields.Line,
				Pattern:     tt.fields.Pattern,
				NumPossible: tt.fields.NumPossible,
			}
			n.Unfold(tt.args.c)
			if !reflect.DeepEqual(n.Pattern, tt.want.pattern) {
				t.Errorf("Unfold().Pattern = %v, want %v", n.Pattern, tt.want)
			}
			if !reflect.DeepEqual(n.Line, tt.want.line) {
				t.Errorf("Unfold().Line = %v, want %v", n.Line, tt.want)
			}
		})
	}
}

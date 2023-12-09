package algos

import (
	"testing"
)

func TestLCMs(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 indexes",
			args: args{
				s: make([]int, 0),
			},
			want: 0,
		},
		{
			name: "1 index",
			args: args{
				s: []int{
					2,
				},
			},
			want: 2,
		},
		{
			name: "2 indexes",
			args: args{
				s: []int{
					2, 3,
				},
			},
			want: 6,
		},
		{
			name: "3 indexes",
			args: args{
				s: []int{
					2, 3, 5,
				},
			},
			want: 30,
		},
		{
			name: "4 indexes",
			args: args{
				s: []int{
					2, 3, 5, 7,
				},
			},
			want: 210,
		},
		{
			name: "4 indexes - 2s",
			args: args{
				s: []int{
					2, 2, 2, 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCMs(tt.args.s); got != tt.want {
				t.Errorf("LCMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCDs(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 indexes",
			args: args{
				s: make([]int, 0),
			},
			want: 0,
		},
		{
			name: "1 index",
			args: args{
				s: []int{
					2,
				},
			},
			want: 2,
		},
		{
			name: "2 indexes",
			args: args{
				s: []int{
					2, 4,
				},
			},
			want: 2,
		},
		{
			name: "3 indexes",
			args: args{
				s: []int{
					2, 4, 8,
				},
			},
			want: 2,
		},
		{
			name: "4 indexes",
			args: args{
				s: []int{
					2, 4, 8, 16,
				},
			},
			want: 2,
		},
		{
			name: "4 indexes - 2s",
			args: args{
				s: []int{
					2, 2, 2, 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDs(tt.args.s); got != tt.want {
				t.Errorf("GCDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	sortTest(Merge, t)
}

func TestQuick(t *testing.T) {
	sortTest(Quick, t)
}

func TestInsert(t *testing.T) {
	sortTest(Insert, t)
}

func TestSelect(t *testing.T) {
	sortTest(Select, t)
}

func TestBubble(t *testing.T) {
	sortTest(Bubble, t)
}

func TestHeap(t *testing.T) {
	sortTest(Heap, t)
}

func sortTest(sorter func([]byte) []byte, t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			"01",
			[]byte{},
			[]byte{},
		},
		{
			"02",
			[]byte{1},
			[]byte{1},
		},
		{
			"03",
			[]byte{3, 1, 5, 4, 2},
			[]byte{1, 2, 3, 4, 5},
		},
		{
			"04",
			[]byte{3, 1, 4, 2},
			[]byte{1, 2, 3, 4},
		},
		{
			"05",
			[]byte{9, 7, 4, 3, 1},
			[]byte{1, 3, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if out := sorter(tt.in); !reflect.DeepEqual(out, tt.out) {
				t.Errorf("sorter() = %v, want %v", out, tt.out)
			}
		})
	}
}

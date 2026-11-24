package main

import (
	"reflect"
	"testing"
)

func TestPaths(t *testing.T) {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	want := []string{"DDRDRR", "DRDDRR"}
	if got := Paths(maze); !reflect.DeepEqual(got, want) {
		t.Errorf("Paths = %v, want %v", got, want)
	}

	// Blocked start or end: no paths.
	if got := Paths([][]int{{0, 1}, {1, 1}}); len(got) != 0 {
		t.Errorf("blocked start should give no paths, got %v", got)
	}
	if got := Paths([][]int{{1, 1}, {1, 0}}); len(got) != 0 {
		t.Errorf("blocked end should give no paths, got %v", got)
	}

	// Single open cell: one empty path.
	if got := Paths([][]int{{1}}); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("1x1 maze: %v", got)
	}

	// Two routes around an open 2x2.
	if got := Paths([][]int{{1, 1}, {1, 1}}); !reflect.DeepEqual(got, []string{"DR", "RD"}) {
		t.Errorf("2x2 maze: %v", got)
	}
}

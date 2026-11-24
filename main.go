package main

import "fmt"

// Paths returns every route from the top-left to the bottom-right of a grid
// (1 = open, 0 = blocked), moving Down, Left, Right or Up without revisiting
// a cell. Paths are strings over DLRU, generated in lexicographic order by
// trying the moves in that order.
func Paths(maze [][]int) []string {
	n := len(maze)
	out := []string{}
	if n == 0 || maze[0][0] == 0 || maze[n-1][len(maze[n-1])-1] == 0 {
		return out
	}
	m := len(maze[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	moves := []struct {
		ch     byte
		dr, dc int
	}{{'D', 1, 0}, {'L', 0, -1}, {'R', 0, 1}, {'U', -1, 0}}
	path := []byte{}
	var walk func(r, c int)
	walk = func(r, c int) {
		if r == n-1 && c == m-1 {
			out = append(out, string(path))
			return
		}
		visited[r][c] = true
		for _, mv := range moves {
			nr, nc := r+mv.dr, c+mv.dc
			if nr < 0 || nr >= n || nc < 0 || nc >= m || maze[nr][nc] == 0 || visited[nr][nc] {
				continue
			}
			path = append(path, mv.ch)
			walk(nr, nc)
			path = path[:len(path)-1]
		}
		visited[r][c] = false
	}
	walk(0, 0)
	return out
}

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	fmt.Println(Paths(maze))
}

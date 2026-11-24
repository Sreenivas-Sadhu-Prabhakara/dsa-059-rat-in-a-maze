# Rat in a Maze

> Striver SDE Sheet 059 · Day 10 Backtracking — DFS in all four directions with a visited grid; trying moves in DLRU order emits paths pre-sorted. Go, stdlib only, tested.

## Problem

A rat starts at the top-left of a grid (1 = open, 0 = wall) and must reach the bottom-right. It may move Down, Left, Right or Up but never revisit a cell. Return every possible route as a string of moves, in lexicographic order.

## Approach

Depth-first backtracking: mark the current cell visited, try each legal neighbor, and unmark on the way out — the unmark is what distinguishes "all paths" from plain reachability DFS, since other routes may pass through this cell. Because up-moves are legal, the visited grid is load-bearing: without it the rat oscillates forever between two cells. Trying moves in alphabetical order (D, L, R, U) makes the output arrive already sorted, a cheaper trick than sorting afterwards. Start or goal blocked is checked up front.

**Complexity:** exponential in grid size (it enumerates paths), O(n·m) state.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).

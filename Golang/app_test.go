package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGrid(t *testing.T) {
	got := makeGrid(5, 6)

	assert.Equal(t, 5, len(got))
	assert.Equal(t, 6, len(got[0]))
	assert.Equal(t, false, got[0][0])
}

func TestPopulateGrid(t *testing.T) {
	grid := makeGrid(5, 6)
	aliveCells := []Point{{0, 0}}
	got := populateGrid(aliveCells, grid)
	assert.Equal(t, true, got[0][0])
	assert.Equal(t, false, got[0][1])
	assert.Equal(t, false, got[1][0])
}

func TestCountAliveCells(t *testing.T) {
	grid := makeGrid(5, 6)
	aliveCells := []Point{{0, 0}, {4, 5}}
	got := populateGrid(aliveCells, grid)
	aliveCellsOnly := getAliveCells(got)
	assert.Equal(t, aliveCells, aliveCellsOnly)
	assert.Equal(t, 2, len(aliveCellsOnly))
}

func TestCloneGrid(t *testing.T) {
	grid := makeGrid(5, 6)
	aliveCells := []Point{{0, 0}, {4, 5}}
	grid = populateGrid(aliveCells, grid)
	newGrid := cloneGrid(grid)
	grid[2][2] = true

	assert.NotEqual(t, getAliveCells(grid), getAliveCells(newGrid))
	assert.Equal(t, false, newGrid[2][2])
}

func TestAlivePopulationAfterFirstGeneration(t *testing.T) {
	aliveCells := []Point{{0, 0}, {4, 5}}
	neighbourhood := initializeNeighbourhood(5, 6, aliveCells)
	aliveCellsInNeighbourhood := getAliveCells(neighbourhood)
	assert.Equal(t, 2, len(aliveCellsInNeighbourhood))
}

func TestGetAliveNeighbourCountNoneExpected(t *testing.T) {
	aliveCells := []Point{{2, 2}}
	neighbourhood := initializeNeighbourhood(6, 6, aliveCells)
	count, _ := decideNeighbours(neighbourhood, Point{2, 2}, map[Point]int{})
	assert.Equal(t, 0, count)
}

func TestGetAliveNeighbourCountOneExpected(t *testing.T) {
	aliveCells := []Point{{2, 2}, {2, 3}}
	neighbourhood := initializeNeighbourhood(6, 6, aliveCells)
	count, _ := decideNeighbours(neighbourhood, Point{2, 2}, map[Point]int{})
	assert.Equal(t, 1, count)
}

func TestGetAliveNeighbourCountOneExpectedDiagonal(t *testing.T) {
	aliveCells := []Point{{2, 2}, {3, 3}}
	neighbourhood := initializeNeighbourhood(6, 6, aliveCells)
	count, _ := decideNeighbours(neighbourhood, Point{2, 2}, map[Point]int{})
	assert.Equal(t, 1, count)
}

func TestDeadPopulationAfterGeneration(t *testing.T) {
	aliveCells := []Point{{0, 0}, {4, 5}}
	neighbourhood := initializeNeighbourhood(5, 6, aliveCells)
	var aliveCellsCount int
	neighbourhood, aliveCellsCount = tick(neighbourhood)
	assert.Equal(t, 0, aliveCellsCount)
}

func TestStablePopulationAfterGeneration(t *testing.T) {
	aliveCells := []Point{{1, 1}, {1, 2}, {2, 1}, {2, 2}}
	neighbourhood := initializeNeighbourhood(5, 6, aliveCells)
	var aliveCellsCount int
	neighbourhood, aliveCellsCount = tick(neighbourhood)
	assert.Equal(t, 4, aliveCellsCount)
}

func TestReanimatePopulationAfterGeneration(t *testing.T) {
	aliveCells := []Point{{0, 0}, {0, 1}, {1, 0}}
	neighbourhood := initializeNeighbourhood(5, 6, aliveCells)
	var aliveCellsCount int
	neighbourhood, aliveCellsCount = tick(neighbourhood)
	assert.Equal(t, 4, aliveCellsCount)
	assert.Equal(t, Point{1, 1}, getAliveCells(neighbourhood)[3])
}

// func TestGridCreate(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		m    int
// 		n    int
// 		x    int
// 		y    int
// 		want int
// 	}{
// 		{"0 Neighbours", 5, 5, 2, 2, 0}
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := main.makeGrid(tt.m, tt.n)tt.gb.Neighbours(tt.x, tt.y); got != tt.want {
// 				main.makeGrid(5,5)
// 				tt.gb.Print()
// 				t.Errorf("GameBoard.Neighbours() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

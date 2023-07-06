package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]bool
type Point struct {
	X int
	Y int
}

func main() {
	grid := initializeNeighbourhood(10, 10, []Point{
		{0, 0}, {0, 1}, {1, 0}, {2, 3}, {3, 2}, {3, 3},
		{7, 0}, {7, 1}, {7, 2},
	})
	aliveCount := len(getAliveCells(grid))
	for aliveCount > 0 {
		printGrid(grid)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		grid, aliveCount = tick(grid)
	}
	print("THE END!")
}

func printGrid(grid Grid) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}

func makeGrid(x, y int) Grid {
	grid := make(Grid, x)
	for i := range grid {
		grid[i] = make([]bool, y)
	}
	return grid
}

func populateGrid(liveCells []Point, grid Grid) Grid {
	for _, element := range liveCells {
		grid[element.X][element.Y] = true
	}
	return grid
}

func initializeNeighbourhood(width, depth int, liveCells []Point) Grid {
	grid := makeGrid(width, depth)
	return populateGrid(liveCells, grid)
}

func getAliveCells(grid Grid) []Point {
	var aliveCells []Point
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell {
				aliveCells = append(aliveCells, Point{rowIndex, colIndex})
			}
		}
	}
	return aliveCells
}

func cloneGrid(grid Grid) Grid {
	newGrid := makeGrid(len(grid), len(grid[0]))
	return populateGrid(getAliveCells(grid), newGrid)
}

func decideNeighbours(grid Grid, point Point, deadMap map[Point]int) (int, map[Point]int) {
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			row := point.X + r
			col := point.Y + c
			withinBounds := row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
			isNotSelfPoint := point.X != row || point.Y != col
			if withinBounds && isNotSelfPoint {
				if grid[row][col] {
					count++
				} else {
					deadMap[Point{row, col}] += 1
				}
			}
		}
	}
	return count, deadMap
}

func tick(grid Grid) (Grid, int) {
	aliveCells := getAliveCells(grid)
	aliveCellsNextTick := []Point{}
	deadWithAliveNeighbours := map[Point]int{}
	for _, cell := range aliveCells {
		var aliveNeighboursCount int
		aliveNeighboursCount, deadWithAliveNeighbours = decideNeighbours(grid, cell, deadWithAliveNeighbours)
		if aliveNeighboursCount >= 2 && aliveNeighboursCount <= 3 {
			aliveCellsNextTick = append(aliveCellsNextTick, cell)
		}
	}
	for key, deadCell := range deadWithAliveNeighbours {
		if deadCell == 3 {
			aliveCellsNextTick = append(aliveCellsNextTick, key)
		}
	}
	return initializeNeighbourhood(len(grid), len(grid[0]), aliveCellsNextTick), len(aliveCellsNextTick)
}

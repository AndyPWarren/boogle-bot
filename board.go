package main

// Pos position in the grid
type Pos struct {
	X int
	Y int
}

// GetSurroundingPos given a position in the grid and a row size return all position surrounding positions (up, down, left, right and diagonally)
func GetSurroundingPos(p Pos, rowSize int) []Pos {
	p1 := Pos{p.X - 1, p.Y - 1}
	p2 := Pos{p.X - 1, p.Y}
	p3 := Pos{p.X - 1, p.Y + 1}
	p4 := Pos{p.X, p.Y - 1}
	p5 := Pos{p.X, p.Y + 1}
	p6 := Pos{p.X + 1, p.Y - 1}
	p7 := Pos{p.X + 1, p.Y}
	p8 := Pos{p.X + 1, p.Y + 1}
	allPs := []Pos{p1, p2, p3, p4, p5, p6, p7, p8}
	filteredPs := []Pos{}
	for _, p := range allPs {
		if p.X >= 0 && p.Y >= 0 && p.X <= rowSize-1 && p.Y <= rowSize-1 {
			filteredPs = append(filteredPs, p)
		}
	}
	return filteredPs
}

// CharsToGrid given a string of characters and a grid size, map chars to arrays of rows
func CharsToGrid(chars []string, gridSize int) [][]string {
	rows := [][]string{}
	for i := 0; i < gridSize; i++ {
		row := chars[i*gridSize : i*gridSize+gridSize]
		rows = append(rows, row)
	}
	return rows
}

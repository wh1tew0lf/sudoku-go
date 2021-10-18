package main

import "fmt"

type SudokuField struct {
	Value int
	Values []int
}

func initSudoku(n int) [][]SudokuField {
	sudoku := make([][]SudokuField, n)
	for i:=0; i<n; i++ {
		sudoku[i] = make([]SudokuField, n)
		for j,_ := range sudoku[i] {
			var numb int
			fmt.Scan(&numb)
			if numb > 0 {
				sudoku[i][j].Value = numb;
			} else {
				sudoku[i][j].Value = 0;
				sudoku[i][j].Values = make([]int, n)
				for k:=0; k<n; k++ {
					sudoku[i][j].Values[k] = 1
				}
			}
		}
	}
	
	return sudoku
}

func printSudoku(sudoku [][]SudokuField, size int) {	
	for i:=0; i < size; i++ {
		for j:=0; j < size; j++ {
			if (sudoku[i][j].Value > 0) {
				fmt.Print(" ", sudoku[i][j].Value, "  ")
			} else {
				var unfoundCount = 0
				for k:=0; k < size; k++ {
					if (sudoku[i][j].Values[k] > 0) {
						unfoundCount++
					}
				}
				fmt.Print("[", unfoundCount ,"] ")
			}
				
		}
		fmt.Println()
	}
}

func checkCell(values []int, value int) bool {
	if value > 0 && values[value - 1] > 0 {
		values[value - 1] = 0
		return true
	}
	return false
}

func checkNumberInLine(number int, sudoku [][]SudokuField, line int, size int) int {
	var equalsCount = 0
	var equalsIndex int
	
	for k:=0; k < size; k++ {
		if (sudoku[line][k].Value == number) {
			equalsCount = 0
			break
		}
		
		if (sudoku[line][k].Value > 0) {
			continue
		}
		
		if sudoku[line][k].Values[number-1] > 0 {
			equalsCount++
			equalsIndex = k
		}
	}
	
	if equalsCount == 1 {
		sudoku[line][equalsIndex].Value = number
		return 1
	}
	
	return 0
}

func checkNumberInColumn(number int, sudoku [][]SudokuField, column int, size int) int {
	var equalsCount = 0
	var equalsIndex int
	
	for line:=0; line < size; line++ {
		if (sudoku[line][column].Value == number) {
			equalsCount = 0
			break
		}
		
		if (sudoku[line][column].Value > 0) {
			continue
		}
		
		if sudoku[line][column].Values[number-1] > 0 {
			equalsCount++
			equalsIndex = line
		}
	}
	
	if equalsCount == 1 {
		sudoku[equalsIndex][column].Value = number
		return 1
	}
	
	return 0
}

func checkNumberInSquare(number int, sudoku [][]SudokuField, i int, j int, size int) int {
	var equalsCount = 0
	var equalsColumnIndex int
	var equalsLineIndex int
	
	for k:=i/3*3; k < (i/3+1)*3; k++ {
		for l:=j/3*3; l < (j/3+1)*3; l++ {			
			if (sudoku[k][l].Value == number) {
				return 0
			}
		
			if (sudoku[k][l].Value > 0) {
				continue
			}
			
			if sudoku[k][l].Values[number-1] > 0 {
				equalsCount++
				equalsLineIndex = k
				equalsColumnIndex = l
				
				if equalsCount > 1 {
					return 0
				}
			}
		}
	}
	
	if equalsCount == 1 {
		sudoku[equalsLineIndex][equalsColumnIndex].Value = number
		return 1
	}
	
	return 0
}

func checkLine(sudoku [][]SudokuField, i int, j int, size int) int {
	var changesCount = 0
	
	for k:=0; k < size; k++ {
		if checkCell(sudoku[i][j].Values, sudoku[i][k].Value) {
			changesCount++
		}
	}
	return changesCount
}

func checkColumn(sudoku [][]SudokuField, i int, j int, size int) int {
	var changesCount = 0
	
	for k:=0; k < size; k++ {
		if checkCell(sudoku[i][j].Values, sudoku[k][j].Value) {
			changesCount++
		}
	}
	return changesCount
}

func checkSquare(sudoku [][]SudokuField, i int, j int, size int) int {
	var changesCount = 0
	
	for k:=i/3*3; k < (i/3+1)*3; k++ {
		for l:=j/3*3; l < (j/3+1)*3; l++ {
			if checkCell(sudoku[i][j].Values, sudoku[k][l].Value) {
				changesCount++
			}
		}
	}
	return changesCount
}

func tryExtractValue(sudokuCell SudokuField, size int) SudokuField {
	var valuesCount = 0
	var value = 0
	for k:=0; k<size; k++ {
		if sudokuCell.Values[k] != 0 {
			valuesCount++
			value = k+1
		}
	}
	if valuesCount == 1 {
		sudokuCell.Value = value
	}
	
	return sudokuCell
}

func solveIteration(sudoku [][]SudokuField, size int) int {
	var changesCount = 0
	for i:=0; i < size; i++ {
		for j:=0; j < size; j++ {
			if (sudoku[i][j].Value > 0) {
				continue
			}
			
			changesCount += checkLine(sudoku, i, j, size);
			changesCount += checkColumn(sudoku, i, j, size);
			changesCount += checkSquare(sudoku, i, j, size);
			
			for k:=0; k<size; k++ {
				changesCount += checkNumberInLine(k+1, sudoku, i, size)				
				changesCount += checkNumberInColumn(k+1, sudoku, j, size)			
				changesCount += checkNumberInSquare(k+1, sudoku, i, j, size)
			}
			
			// check which 
			sudoku[i][j] = tryExtractValue(sudoku[i][j], size)			
		}
	}
	return changesCount
}

func main() {
	var n = 9
	var sudoku = initSudoku(n)
	for solveIteration(sudoku, n) > 0 {
		fmt.Println("Iteration...")
	}
	printSudoku(sudoku, n)
}

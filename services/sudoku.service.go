package services

func Sort(sudoku [9][9]int) [][]int {
	var hasilAkhir [][]int
	for i := 0; i < len(sudoku); i++ {
		var hasil []int
		for j := 0; j < len(sudoku); j++ {
			hasil = append(hasil, sudoku[i][j])
		}
		hasilAkhir = append(hasilAkhir, hasil)
	}
	return hasilAkhir
}

func CheckPadaBaris(k int, sudoku [9][9]int, x int) bool {
	for i := 0; i < len(sudoku); i++ {
		if sudoku[x][i] == k {
			return false
		}
	}
	return true
}

func CheckPadaRow(k int, sudoku [9][9]int, y int) bool {
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i][y] == k {
			return false
		}
	}
	return true
}

func CheckPerColumn(k int, sudoku [9][9]int, x int, y int) bool {
	Xpertama := x - (x % 3)
	Ypertama := y - (y % 3)
	for i := Xpertama; i < Xpertama+3; i++ {
		for j := Ypertama; j < Ypertama+3; j++ {
			if sudoku[i][j] == k {
				return false
			}
		}
	}
	return true
}

func IsValid(sudoku *[9][9]int, position int) bool {
	if position == 9*9 {
		return true
	}

	x := position / 9
	y := position % 9

	if sudoku[x][y] != 0 {
		return IsValid(sudoku, position+1)
	}
	for i := 1; i <= 9; i++ {
		if CheckPadaBaris(i, *sudoku, x) && CheckPadaRow(i, *sudoku, y) && CheckPerColumn(i, *sudoku, x, y) {
			sudoku[x][y] = i

			if IsValid(sudoku, position+1) {
				return true
			}
		}
	}
	sudoku[x][y] = 0
	return false
}

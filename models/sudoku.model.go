package models

type Sudoku struct {
	Sudo [9][9]int `json:"data,omitempty"`
}

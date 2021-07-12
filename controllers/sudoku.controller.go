package controllers

import (
	"testidn/models"
	"testidn/services"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool    `json:"success"`
	Data    [][]int `json:"data"`
}

func SudokuAction(c *gin.Context) {
	var req models.Sudoku

	c.ShouldBindJSON(&req)

	result := services.IsValid(&req.Sudo, 0)
	if result {
		c.JSON(200, &Response{
			Success: true,
			Data:    services.Sort(req.Sudo),
		})
	}
}

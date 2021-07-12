package routes

import (
	"testidn/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/answer", controllers.SudokuAction)
	}

}

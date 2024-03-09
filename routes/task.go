package routes

import (
	"github.com/atirach-in/management-system.git/controllers"
	"github.com/gin-gonic/gin"
)

func InitTaskRoute(rg *gin.RouterGroup) {
	rg.GET("/all", controllers.AllTasks)
	rg.POST("/create", controllers.CreateTask)
	rg.PUT("/update", controllers.UpdateTask)
	rg.PUT("/update/status", controllers.UpdateStatusTask)
	rg.GET("/by/:id", controllers.TaskById)
	rg.DELETE("/:id", controllers.DelTask)
}

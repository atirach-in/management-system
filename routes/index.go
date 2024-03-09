package routes

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	api := r.Group("/api")
	{
		task := api.Group("/task")
		task.Use()
		{
			InitTaskRoute(task)
		}
	}
}

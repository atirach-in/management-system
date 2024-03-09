package controllers

import (
	"fmt"
	"log"

	"github.com/atirach-in/management-system.git/configs"
	"github.com/atirach-in/management-system.git/model"
	"github.com/gin-gonic/gin"
)

type reqCreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type reqUpdateTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type reqUpdateStatusTask struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

func AllTasks(ctx *gin.Context) {
	db := configs.ConnectDB()

	results, err := db.Query("SELECT * FROM tasks")
	defer db.Close()

	if err != nil {
		fmt.Println("All task error")
		panic(err.Error())
	}

	var tasksData []model.Task
	for results.Next() {
		var task model.Task
		err := results.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
		)

		if err != nil {
			fmt.Printf("err scan : ", err)
			panic(err.Error())
		}

		tasksData = append(tasksData, task)
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "All tasks",
		"obj":     tasksData,
	})
	return
}

func CreateTask(ctx *gin.Context) {

	var body reqCreateTask
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(422, gin.H{
			"status":  422,
			"success": false,
			"data":    "Invalid data",
		})
	}

	db := configs.ConnectDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO tasks (Title, Description, Status) VALUE(?, ?, ?)", body.Title, body.Description, body.Status)

	if err != nil {
		fmt.Print("\n\nErr insert data: ", err.Error())
		ctx.JSON(500, gin.H{
			"status":  500,
			"success": false,
			"data":    "Create task failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "Create task success",
	})
	return
}

func UpdateTask(ctx *gin.Context) {
	var body reqUpdateTask

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(422, gin.H{
			"status":  422,
			"success": false,
			"data":    "Invalid data",
		})
		return
	}

	var db = configs.ConnectDB()
	defer db.Close()

	_, err := db.Exec("UPDATE tasks SET Title = ?, Description = ? WHERE ID = ?", body.Title, body.Description, body.ID)
	if err != nil {
		log.Fatal(err)

		ctx.JSON(500, gin.H{
			"status":  500,
			"success": false,
			"data":    "Update task failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "Update task success",
	})
	return
}

func UpdateStatusTask(ctx *gin.Context) {
	var body reqUpdateStatusTask

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(422, gin.H{
			"status":  422,
			"success": false,
			"data":    "Invalid data",
		})
		return
	}

	var db = configs.ConnectDB()
	defer db.Close()

	_, err := db.Exec("UPDATE tasks SET Status = ? WHERE ID = ?", body.Status, body.ID)
	if err != nil {
		log.Fatal(err)

		ctx.JSON(500, gin.H{
			"status":  500,
			"success": false,
			"data":    "Update task failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "Update status task success",
	})
	return
}

func TaskById(ctx *gin.Context) {

	id := ctx.Param("id")

	db := configs.ConnectDB()

	var task model.Task

	db.QueryRow("SELECT * FROM tasks WHERE ID = ?", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	defer db.Close()

	if task.ID == 0 {
		ctx.JSON(200, gin.H{
			"status":  200,
			"success": false,
			"data":    "Not found data",
			"obj":     task,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "Task by id",
		"obj":     task,
	})
	return
}

func DelTask(ctx *gin.Context) {
	id := ctx.Param("id")

	db := configs.ConnectDB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM tasks WHERE ID = ?", id)
	if err != nil {
		log.Fatal(err)

		ctx.JSON(500, gin.H{
			"status":  500,
			"success": false,
			"data":    "Delete task failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"data":    "Delete task success",
	})
	return
}

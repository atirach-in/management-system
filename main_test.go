package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atirach-in/management-system.git/configs"
	"github.com/atirach-in/management-system.git/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	r = gin.New()

	if err := configs.InitDotEnv(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	configs.TestConnectDB()

	// Set up a test route with your handler
	r.GET("/all", controllers.AllTasks)
	r.POST("/create", controllers.CreateTask)
	r.PUT("/update", controllers.UpdateTask)
	r.PUT("/update/status", controllers.UpdateStatusTask)
	r.GET("/by/:id", controllers.TaskById)
	r.DELETE("/:id", controllers.DelTask)
}

func TestAllTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("TestAllTasks: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateTask(t *testing.T) {

	requestBody := []byte(`{
		"title": "Python",
		"description": "Learn python",
		"status": 0
	}`)

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", "/create", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("TestCreateTask: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateTask(t *testing.T) {

	requestBody := []byte(`{
		"id": 1,
		"title": "Python",
		"description": "Learn python"
	}`)

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", "/update", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("TestUpdateTask: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateStatusTask(t *testing.T) {

	requestBody := []byte(`{
		"id": 4,
		"status": 0
	}`)

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", "/update", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("UpdateStatusTask: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestTaskById(t *testing.T) {
	req, err := http.NewRequest("GET", "/by/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("TestTaskById: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDelTask(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("TestDelTask: ", rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

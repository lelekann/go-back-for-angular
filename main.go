package main

import (
	"net/http"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type task struct {
	ID		string	`json: "id"`
	Task	string	`json: "task"`
	Time 	string	`json: "time"`
	Active	bool	`json: "active"`
}

var tasks = []task{
	{ID: "1", Task: "Create Angular App", Time: "12:20 AM", Active: true},
	{ID: "2", Task: "Create back for Angular App", Time: "12:40 AM", Active: true},
	{ID: "3", Task: "Learn Angular", Time: "12:20 PM", Active: false},
	{ID: "4", Task: "Learn Golang", Time: "12:20 AM", Active: true},
}

func getTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

func addTask(context *gin.Context) {
	var newTask task
	
	if err := context.BindJSON(&newTask); err !=nil {
		return
	}

	tasks = append(tasks, newTask)

	context.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskById(id string) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("Task not found")
}

func updateTask(context *gin.Context){
	id := context.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	task.Active = !task.Active

	context.IndentedJSON(http.StatusOK, task)
}

func deleteTask (context *gin.Context) {
	id := context.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted!"})
}

func main () {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
	router.GET("/tasks", getTasks)
	router.POST("/tasks", addTask)
	router.GET("/tasks/:id", updateTask)
	router.DELETE("tasks/:id", deleteTask)
	router.Run("localhost:9090")
}
package controllers

import (
	"net/http"
	"time"

	"headfirstgo/crud/crud_app/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assigned_to"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assigned_to"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

//Get /tasks
//Get all tasks

func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

//Post /tasks
//Create new task

func CreateTask(c *gin.Context) {
	//validate input

	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-07-06"
	deadline, _ := time.Parse(date, input.Deadline)

	//Create task

	task := models.Task{AssignedTo: input.AssignedTo, Task: input.Task, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

//Get /tasks:id
///Find a task

func FindTask(c *gin.Context) {
	//get model if exists
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

//Patch /tasks:id
//Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var task models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//Validate input

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-07-06"
	deadline, _ := time.Parse(date, input.Deadline)

	var updateInput models.Task
	updateInput.Deadline = deadline
	updateInput.AssignedTo = input.AssignedTo
	updateInput.Task = input.Task

	db.Model(&task).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

//Delete /tasks:id
//Delete a task
func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	///get model if exists
	var list models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})

}

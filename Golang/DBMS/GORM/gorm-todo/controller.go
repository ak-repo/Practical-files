package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func TodoCreate(c *gin.Context) {

	//get data from req body
	var body struct {
		Content string
		Done    bool
	}
	c.Bind(&body)

	//create a todo
	todo := Todo{Content: body.Content, Done: body.Done}
	result := DB.Create(&todo)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "not posted in DB"})
		return
	}
	log.Println("create result: ", result)
	//return it
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodoIndex(c *gin.Context) {
	// get all todos
	var todos []Todo
	DB.Find(&todos)

	// return todos in response
	c.JSON(200, gin.H{"todos": todos})
}

func TodosShow(c *gin.Context) {
	//get id from URL params
	id := c.Param("id")

	//get a single todo
	var todo Todo
	DB.First(&todo, id)

	//return todo in response
	c.JSON(200, gin.H{"todo": todo})

}

func TodoUpdate(c *gin.Context) {
	id := c.Param("id")

	//get the data from req body
	var body struct {
		Content string
		Done    bool
	}
	c.Bind(&body)

	//Get a single todo that we what to update\
	var todo Todo
	DB.First(&todo, id)

	DB.Model(&todo).Updates(Todo{Content: body.Content, Done: body.Done})

	c.JSON(200, gin.H{"todo": todo})
}

func TodoDelete(c *gin.Context) {
	id := c.Param("id")

	//delete the todo
	DB.Delete(&Todo{}, id)

	//return response
	c.JSON(200, gin.H{"message": "todo removed"})

}

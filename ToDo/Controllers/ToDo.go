package Controllers

import (
	"fmt"
	"net/http"
	"ToDo/Models"
	"github.com/gin-gonic/gin"
)

func CreateATodo(c *gin.Context) {
	var todo Models.ToDo
	c.BindJSON(&todo)
	// db query here
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllTodos(c *gin.Context) {
	var todo []Models.ToDo
	rows, err := Models.GetAllTodoObjects(&todo)
	if err !=  nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, rows)
	}
}


func GetTodobyID(c *gin.Context) {
	var todo Models.ToDo
	id := c.Params.ByName("id")
	fmt.Println(id)
	todos, err := Models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}


func UpdateTodo(c *gin.Context) {
	var todo Models.ToDo
	id := c.Params.ByName("id")
	_, err := Models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&todo)
	_, err = Models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id: "+id: "updated"})
	}
}

func DeleteTodo(c *gin.Context) {
	var todo Models.ToDo
	id := c.Params.ByName("id")
	_, err := Models.DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id: "+id: "deleted"})
	}
}
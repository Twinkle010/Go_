package Routes

import (
	"ToDo/Controllers"
	"github.com/gin-gonic/gin"
)


func SetUpRouter() *gin.Engine {
	// intilize gin default engine that comes with loggers
	r := gin.Default()
	// group routes using group func
	v1 := r.Group("/v1")
	{
		// route to get existing todo's
		v1.GET("todo", Controllers.GetAllTodos)
		// to create todo tasks
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetTodobyID)
		v1.PUT("todo/:id", Controllers.UpdateTodo)
		v1.DELETE("todo/:id", Controllers.DeleteTodo)
	}
	return r 
}
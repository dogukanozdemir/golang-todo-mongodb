package main

import (
	"net/http"

	controller "github.com/dogukanozdemir/go-todo-mongo/controllers"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")

	// todo
	router.GET("/", index)
	router.GET("/todos/:userid", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)
	router.POST("/todo/:userid", controller.AddTodo)
	router.DELETE("/todo/:userid/:id", controller.DeleteTodo)
	router.DELETE("/todos/:userid", controller.ClearAll)
	router.PUT("/todo", controller.UpdateTodo)

	// user
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/todo", controller.Todo)

	router.Run(":8080")

}

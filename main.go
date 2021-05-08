package main

import (
	"github.com/gin-gonic/gin"

	"golang_bookstore/controllers"
	"golang_bookstore/models"
)

func main()  {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}
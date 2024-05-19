package main

import (
	"BookWebApi/db"
	"BookWebApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response map[string]interface{}

func main() {
	db.Init()

	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {
		limitParam := context.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid limit parameter",
			})
			return
		}

		result, err := models.GetBooks(limit)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(http.StatusOK, Response{
			"message": "All books in the database",
			"books":   result,
		})
	})

	app.GET("/books/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid ID",
			})
			return
		}

		book, err := models.GetBookById(id)
		if err != nil {
			context.JSON(http.StatusNotFound, Response{
				"message": "Book not found",
			})
			return
		}

		context.JSON(http.StatusOK, Response{
			"message": "Book details",
			"book":    book,
		})
	})

	app.POST("/books", func(context *gin.Context) {
		var bookObject models.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid object",
			})
			return
		}

		err = bookObject.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(http.StatusOK, Response{
			"message": "Book created successfully",
			"object":  bookObject,
		})
	})

	app.PUT("/books/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid ID",
			})
			return
		}

		var bookObject models.Book
		err = context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid object",
			})
			return
		}

		bookObject.Id = id
		err = bookObject.Update()
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Cannot update book object",
			})
			return
		}

		context.JSON(http.StatusOK, Response{
			"message": "Book updated successfully",
			"object":  bookObject,
		})
	})

	app.DELETE("/books/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Invalid ID",
			})
			return
		}

		err = models.DeleteBookById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, Response{
				"message": "Cannot delete book object",
			})
			return
		}

		context.JSON(http.StatusOK, Response{
			"message": "Book deleted successfully",
		})
	})

	err := app.Run(":8080")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}

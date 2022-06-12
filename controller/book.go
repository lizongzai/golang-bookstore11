package controller

import (
	"bookstore/dao/mysql"
	"bookstore/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	http://localhost:8080/api/v1/book
{
    "name": "the Hobbit 2",
    "author": "JJR Tolkien"
}
*/

func CreateBookHandler(c *gin.Context) {
	book := new(model.Book)
	if err := c.ShouldBind(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	mysql.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

/*
	http://localhost:8080/api/v1/book
*/
func GetAllBookHandler(c *gin.Context) {
	book := []model.Book{}
	mysql.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

/*
	http://localhost:8080/api/v1/book/3/
*/
func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	bookIds, _ := strconv.ParseInt(bookId, 0, 0)
	book := model.Book{Id: bookIds}
	mysql.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

/*
	http://localhost:8080/api/v1/book
{
    "id": 5,
    "name": "the Hobbit III",
    "author": "JJR Tolkien"
}
*/
func UpdateBookDetailHandler(c *gin.Context) {
	book := new(model.Book)
	if err := c.ShouldBind(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	oldBook := &model.Book{Id: book.Id}
	var newBook model.Book
	if book.Name != "" {
		newBook.Name = book.Name
	}

	if book.Author != "" {
		newBook.Author = book.Author
	}

	mysql.DB.Model(&oldBook).Updates(&newBook)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

/*
	http://localhost:8080/api/v1/book/4/
*/
func DeleteBookHandler(c *gin.Context) {
	bookId := c.Param("id")
	bookIds, _ := strconv.ParseInt(bookId, 0, 0)
	mysql.DB.Select("User").Delete(&model.Book{Id: bookIds})
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

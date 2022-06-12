package controller

import (
	"bookstore/dao/mysql"
	"bookstore/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 书籍入库
func CreateBookHandler(c *gin.Context) {
	book := new(model.Book)
	if err := c.ShouldBind(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	mysql.DB.Create(book)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 查询所有记录
func GetAllBookHandler(c *gin.Context) {
	book := []model.Book{}
	mysql.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

// 查询单个记录
func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	bookIds, _ := strconv.ParseInt(bookId, 0, 0)
	book := model.Book{Id: bookIds}
	mysql.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

// 修改书籍记录
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

// 删除书籍记录
func DeleteBookHandler(c *gin.Context) {
	bookId := c.Param("id")
	bookIds, _ := strconv.ParseInt(bookId, 0, 0)
	mysql.DB.Select("User").Delete(&model.Book{Id: bookIds})
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

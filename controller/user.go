package controller

import (
	"bookstore/dao/mysql"
	"bookstore/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 用户注册
func RegisterHandler(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	mysql.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 用户登录
func LoginHandler(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	u := model.User{Name: user.Name, Password: user.Password}
	if row := mysql.DB.Where(&u).First(&u).Row(); row == nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMsg": "The user and passowrd do not exist"})
		return
	}

	token := uuid.New().String()
	mysql.DB.Model(u).Update("token", token)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "token": token})
}

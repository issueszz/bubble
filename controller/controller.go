package controller

import (
	"bubble/common"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler 主页
func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}
func GetToDo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail("无效id"))
			return
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
		} else {
			c.JSON(http.StatusOK, common.OK.WithData(todo))
		}
	}
}
// CreateTodo 生成某个待办事项并写入数据库
func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.ToDo
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
			return
		}
		if err := models.CreateTodo(&todo); err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
		} else {
			c.JSON(http.StatusOK, common.OK)
		}
	}
}

// GetAllTodo 从数据库中查询所有待办事项并返回前端
func GetAllTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todolist, err := models.GetAllTodo()
		if err != nil {
			c.JSON(http.StatusOK, common.Err.WithDetail(err.Error()))
		} else {
			c.JSON(http.StatusOK, common.OK.WithData(todolist))
		}
	}
}

// UpdateTodo 更新某个待办事项
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail("无效id"))
			return
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
			return
		}


		if err = c.BindJSON(todo); err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
			return
		}

		if err = models.UpdateTodo(todo); err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
		} else {
			c.JSON(http.StatusOK, common.OK)
		}
	}
}

// DeleteTodo 删除某个待办事项
func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail("无效id"))
			return
		}

		if err := models.DeleteTodo(id); err != nil {
			c.JSON(http.StatusInternalServerError, common.Err.WithDetail(err.Error()))
		} else {
			c.JSON(http.StatusOK, common.OK.WithDetail("delete success"))
		}
	}
}







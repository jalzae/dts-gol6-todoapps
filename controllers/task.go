package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest/config"
	"rest/models"
)

type taskRepo struct {
	Db *gorm.DB
}

func TaskControll() *taskRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.Task{})
	return &taskRepo{Db: db}
}
func (repository *taskRepo) Createtask(c *gin.Context) {
	var user models.Task

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Createtask(repository.Db, &user)
	c.JSON(http.StatusOK, gin.H{"data": user})

}
func (repository *taskRepo) Gettasks(c *gin.Context) {
	var user []models.Task
	err := models.Gettasks(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (repository *taskRepo) Gettask(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user models.Task
	err := models.Gettask(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (repository *taskRepo) Updatetask(c *gin.Context) {
	var user models.Task
	id, _ := c.Params.Get("id")
	err := models.Gettask(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.Updatetask(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (repository *taskRepo) Deletetask(c *gin.Context) {
	var user models.Task
	id, _ := c.Params.Get("id")
	err := models.Deletetask(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

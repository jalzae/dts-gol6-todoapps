package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest/config"
	"rest/models"
)

type employeRepo struct {
	Db *gorm.DB
}

func EmployeControll() *employeRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.Employe{})
	return &employeRepo{Db: db}
}
func (repository *employeRepo) Createemploye(c *gin.Context) {
	var user models.Employe

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Createemploye(repository.Db, &user)
	c.JSON(http.StatusOK, gin.H{"data": user})

}
func (repository *employeRepo) Getemployes(c *gin.Context) {
	var user []models.Employe
	err := models.Getemployes(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (repository *employeRepo) Getemploye(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user models.Employe
	err := models.Getemploye(repository.Db, &user, id)
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
func (repository *employeRepo) Updateemploye(c *gin.Context) {
	var user models.Employe
	id, _ := c.Params.Get("id")
	err := models.Getemploye(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.Updateemploye(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (repository *employeRepo) Deleteemploye(c *gin.Context) {
	var user models.Employe
	id, _ := c.Params.Get("id")
	err := models.Deleteemploye(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employe deleted successfully"})
}

package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexOption(c *gin.Context) {
	var option []models.Option
	config.DB.Find(&option)
	c.JSON(http.StatusOK, option)
}

func CreateOption(c *gin.Context) {
	var option models.Option
	if err := c.BindJSON(&option); err != nil {
		panic(err)
	}
	if err := config.DB.Create(&option).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func ShowOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
	} else {
		c.JSON(http.StatusOK, option)
	}
}

func UpdateOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := c.BindJSON(&option); err != nil {
		panic(err)
	}
	if err := config.DB.Updates(&option).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func DeleteOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&option).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
package controllers

import (
	"github.com/dmirubtsov/swag-example/models"
	"github.com/gin-gonic/gin"
)

type userControllers struct{}

var User userControllers

// @Summary Update current user
// @ID update-user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /v1/user [put]
func (*userControllers) Update(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user, err := models.User.Update(user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "updating error",
		})
		return
	}
	c.JSON(200, user)
}

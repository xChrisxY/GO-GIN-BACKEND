package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-backend/configs"
	"github.com/go-backend/models"
)

type PersonRequestBody struct {

	Name string `json:"name"`
	Address string `json:"address"`
	Phone uint `json:"phone"`

}

func GetPersons(c *gin.Context) {

	var persons []models.Person

	configs.DB.Find(&persons)
	c.JSON(http.StatusOK, &persons)

}

func CreatePerson(c *gin.Context) {

	body := PersonRequestBody{}

	c.BindJSON(&body)

	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Create(&person)

	if result.Error != nil {

		c.JSON(500, gin.H{"Error":"Failed to insert"})
		return

	}

	c.JSON(http.StatusOK, &person)

}

func GetPersonById(c *gin.Context) {

	id := c.Param("id")
	var person models.Person

	configs.DB.First(&person, id)

	// De prueba
	if  person.ID != 0 {

		c.JSON(http.StatusOK, &person)
		return

	} 

	c.JSON(400, gin.H{"message":"Person not found"})
 	
}

func UpdatePerson(c *gin.Context) {

	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)

	body := PersonRequestBody{}

	c.BindJSON(&body)
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Model(&person).Updates(data)

	if result.Error != nil {

		c.JSON(500, gin.H{"Error": true, "message":"Failed to update"})
		return

	}

	c.JSON(http.StatusOK, &person)

}

func DeletePerson(c *gin.Context) {

	id := c.Param("id")
	var person models.Person
	configs.DB.Delete(&person, id)

	c.JSON(200, gin.H{"deleted":true})

}


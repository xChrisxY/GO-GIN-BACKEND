package main

import (
	"github.com/go-backend/configs"
	"github.com/go-backend/models"
)

func init() {

	configs.ConnectToDB()

}

func main() {

	configs.DB.AutoMigrate(&models.Person{})

}


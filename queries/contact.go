package queries

import (
	"api_olshop/database"
	"api_olshop/dtos"
	"api_olshop/models"
)

// CreateContact to databse
func CreateContact(name string, email string, phone string, address string) dtos.Response {
	db := database.ConnectToDB()
	operationResult := db.Create(&models.Contact{Name: name, Email: email, Phone: phone, Address: address})
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return dtos.Response{Success: true, Data: operationResult}
}

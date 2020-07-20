package queries

import (
	"api_olshop/database"
	"api_olshop/models"
)

// GetRoles get token by token string
func GetRoles(status int) []models.Role {
	var roles []models.Role
	db := database.ConnectToDB()
	db.Where("status = ? AND deleted_at IS NULL", status).Find(&roles)
	return roles
}

// Register new user
func Register(Name string, Address string, RoleID string, Phone string, Email string) models.User {
	db := database.ConnectToDB()
	var result = &models.User{Name: Name, Address: Address, RoleID: RoleID, Phone: Phone, Email: Email, Status: "1"}
	db.Create(&result)
	return *result
}

package repositories

import (
	"api_olshop/models"

	"github.com/jinzhu/gorm"
)

// ContactRepository as struc
type ContactRepository struct {
	db *gorm.DB
}

// NewContactRepository return rep
func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

// Save responds
func (r *ContactRepository) Save(contact *models.Contact) RepositoryResult {
	err := r.db.Save(contact).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: contact}

}

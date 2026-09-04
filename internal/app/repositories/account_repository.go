package repositories

import (
	"gorm.io/gorm"
	"internal/app/models"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Placeholder for actual database connection
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetAccount(id string) (*models.Account, error) {
	var account models.Account
	if err := r.db.First(&account, id).Error; err!= nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) UpdateAccount(id string, account *models.Account) error {
	return r.db.Model(&models.Account{}).Where("id =?", id).Updates(account).Error
}

func (r *AccountRepository) DeleteAccount(id string) error {
	return r.db.Delete(&models.Account{}, id).Error
}
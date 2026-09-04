package services

import (
	"internal/app/models"
	"internal/app/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{
		accountRepository: repositories.NewAccountRepository(),
	}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.accountRepository.CreateAccount(account)
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	return s.accountRepository.GetAccount(id)
}

func (s *AccountService) UpdateAccount(id string, account *models.Account) error {
	return s.accountRepository.UpdateAccount(id, account)
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.accountRepository.DeleteAccount(id)
}
package dao

import (
	"demoProject/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExist(userId string) (bool, error) {
	var account model.Account
	err := a.db.Where("user_id = ?", userId).First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		fmt.Printf("AccountDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (a *AccountDao) Create(account model.Account) (bool, error) {
	if err := a.db.Create(&account).Error; err != nil {
		fmt.Printf("AccountDao Create = [%v]", err)
		return false, err
	}
	return true, nil
}

func (a *AccountDao) FindByUserId(userId string) (*model.Account, error) {
	var account model.Account
	if err := a.db.Where("user_id = ?", userId).First(&account).Error; err != nil {
		fmt.Printf("AccountDao FindByUserId = [%v]", err)
		return nil, err
	}
	return &account, nil
}

package dao

import (
	"demoProject/internal/model"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

func (c *ContentDao) Create(detail model.ContentDetail) error {
	if err := c.db.Create(&detail).Error; err != nil {
		log.Printf("content create error = %v", err)
		return err
	}

	return nil
}

func (c *ContentDao) Update(id int, detail model.ContentDetail) error {
	if err := c.db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		log.Printf("content update error = %v", err)
		return err
	}

	return nil
}

func (c *ContentDao) IsExist(contentId int) (bool, error) {
	var contentDetail model.ContentDetail
	err := c.db.Where("id = ?", contentId).First(&contentDetail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		log.Printf("ContentDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

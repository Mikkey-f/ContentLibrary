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

func (c *ContentDao) Delete(id int) error {
	if err := c.db.Where("id = ?", id).Delete(&model.ContentDetail{}).Error; err != nil {
		log.Printf("content delete error = %v", err)
		return err
	}

	return nil
}

type ContentSelectReq struct {
	Id       int `json:"id"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (c *ContentDao) Select(params *ContentSelectReq) ([]*model.ContentDetail, int64, error) {
	//var content model.ContentDetail
	//if err := c.db.Where("id = ?", id).First(&content).Error; err != nil {
	//	log.Printf("content select error = %v", err)
	//	return nil, err
	//}
	query := c.db.Model(&model.ContentDetail{})
	if params.Id != 0 {
		query = query.Where("id = ?", params.Id)
	}
	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}

	offset := (page - 1) * pageSize
	var data []*model.ContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil
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

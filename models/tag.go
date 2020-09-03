package models

import "github.com/jinzhu/gorm"

type Tag struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string `gorm:"column:title;size:255;" `
	Articles []Article `gorm:"many2many:article_tags"`
	gorm.Model
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) Add(title string) error{
	return Db.Create(&Tag{Title:title}).Error
}

// 根據名稱去檢查標籤是否存在
func (t *Tag) ExistByName(title string) (bool, error){
	var tag Tag
	if err := Db.Select("id").Where("title = ? ", title).First(&tag).Error; err != nil {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// 取得標籤
func (t *Tag) GetTags() ([]Tag, error) {
	var tags []Tag
	if err := Db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// 根據編號取得標籤
func (t *Tag) GetById (id int) (*Tag, error) {
	var tag Tag
	if err := Db.Where("id = ?" , id).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// 刪除標籤
func (t *Tag) DeleteById(id int) (bool, error) {
	if err := Db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// 更新標籤資訊
func (t *Tag) UpdateById(id int, data interface{}) (bool, error) {
	if err := Db.Model(&Tag{}).Where("id = ?", id).Update(data).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (t *Tag) GetByIds(id []string) ([]Tag, error) {
	var tag []Tag
	if err := Db.Where("id in (?)", id).Find(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
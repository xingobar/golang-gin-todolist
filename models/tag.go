package models

import "github.com/jinzhu/gorm"

type Tag struct {
	ID int `gorm:"primaryKey";"column:id;autoIncrement"`
	Title string `gorm: "column:title";"size:255" `
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
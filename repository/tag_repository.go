package repository

import (
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
)

type tagRepository struct {

}

func NewTagRepository() interfaces.ITagRepository {
	return &tagRepository{}
}

func (repository *tagRepository) Add(title string) (bool, error) {
	 if err := models.Db.Create(&models.Tag{Title:title}).Error; err != nil {
	 	return false, err
	 }
	 return true, nil
}

func (repository *tagRepository) ExistByName(title string) (bool, error) {
	var tag models.Tag
	if err := models.Db.Select("id").Where("title = ? ", title).First(&tag).Error; err != nil {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (repository *tagRepository) GetTags() ([]models.Tag, error) {
	var tags []models.Tag
	if err := models.Db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (repository *tagRepository) GetById(id int) (*models.Tag, error) {
	var tag models.Tag
	if err := models.Db.Where("id = ?" , id).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (repository *tagRepository) DeleteById(id int) (bool, error) {
	if err := models.Db.Where("id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository *tagRepository) UpdateById(id int, tag models.Tag) (bool, error) {
	if err := models.Db.Model(&models.Tag{}).Where("id = ?", id).Update(tag).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository *tagRepository) GetByIds(id []string) ([]models.Tag, error) {
	var tag []models.Tag
	if err := models.Db.Where("id in (?)", id).Find(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

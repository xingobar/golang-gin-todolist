package tag_service

import (
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
	"golang-gin-todolist/repository"
)

type Tag struct {
	Id int
	Title string
}

type TagService struct {
	tagRepository interfaces.ITagRepository
}

func NewTagService() *TagService{
	return &TagService{
		tagRepository: repository.NewTagRepository(),
	}
}

// 新增標籤
func (s *TagService) CreateTag(title string) bool{
	_, err := s.tagRepository.Add(title)
	if err != nil {
		return false
	}
	return true
}

// 判斷標籤名稱是否存在
func (s *TagService) ExistByName(title string) bool{
	ok, err := s.tagRepository.ExistByName(title)
	if err != nil {
		return false
	}

	if !ok {
		return false
	}
	return true
}

func (s *TagService) GetTags() []models.Tag {

	tags, err := s.tagRepository.GetTags()

	if err != nil {
		return nil
	}
	return tags
}

func (s *TagService) GetById(id int) *models.Tag {
	tag, err := s.tagRepository.GetById(id)
	if err != nil {
		return nil
	}
	return tag
}

func (s *TagService) DeleteById(id int) bool {
	_, err := s.tagRepository.DeleteById(id)
	if err != nil {
		return false
	}
	return true
}

// 更新標籤名稱
func (s *TagService) UpdateById(id int, title string) bool {
	_, err := s.tagRepository.GetById(id)
	if err != nil {
		return false
	}
	if _, err := s.tagRepository.UpdateById(id, models.Tag{Title:title}); err != nil {
		return false
	}
	return true
}

func (s *TagService) GetByIds(id []string) []models.Tag{
	tags, err := s.tagRepository.GetByIds(id)
	if err != nil {
		return nil
	}
	return tags
}
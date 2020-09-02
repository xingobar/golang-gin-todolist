package tag_service

import "golang-gin-todolist/models"

type Tag struct {
	Id int
	Title string
}

type TagService struct {
	tag *models.Tag
}

func NewTagService() *TagService{
	return &TagService{
		tag: &models.Tag{},
	}
}

// 新增標籤
func (s *TagService) CreateTag(title string) (error){
	err := s.tag.Add(title)
	if err != nil {
		return err
	}
	return nil
}

// 判斷標籤名稱是否存在
func (s *TagService) ExistByName(title string) bool{
	exists, err := s.tag.ExistByName(title)
	if err != nil {
		return false
	}
	return exists
}

func (s *TagService) GetTags() []models.Tag {
	tags, err := s.tag.GetTags()
	if err != nil {
		return nil
	}
	return tags
}

func (s *TagService) GetById(id int) *models.Tag {
	tag, err := s.tag.GetById(id)
	if err != nil {
		return nil
	}
	return tag
}

func (s *TagService) DeleteById(id int) bool {
	_, err := s.tag.DeleteById(id)
	if err != nil {
		return false
	}
	return true
}

// 更新標籤名稱
func (s *TagService) UpdateById(id int, title string) bool {
	_, err := s.tag.GetById(id)
	if err != nil {
		return false
	}
	var tag models.Tag
	tag.Title = title
	if _, err := s.tag.UpdateById(id, tag); err != nil {
		return false
	}
	return true
}

func (s *TagService) GetByIds(id []string) []models.Tag{
	tags, err := s.tag.GetByIds(id)
	if err != nil {
		return nil
	}
	return tags
}
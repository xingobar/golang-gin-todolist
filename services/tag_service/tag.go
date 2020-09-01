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

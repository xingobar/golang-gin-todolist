package tag_service

import "golang-gin-todolist/models"

type Tag struct {
	Id int
	Title string
}

func (t *Tag) AddTag() error {
	return models.AddTag(t.Title)
}
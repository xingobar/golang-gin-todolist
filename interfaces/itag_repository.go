package interfaces

import "golang-gin-todolist/models"

type ITagRepository interface {
	// 新增
	Add(title string) (bool, error)

	// 確認標籤是否存在
	ExistByName(title string) (bool, error)

	// 取得所有標籤
	GetTags() ([]models.Tag, error)

	// 根據編號取得標籤
	GetById(id int) (*models.Tag, error)

	// 根據編號刪除
	DeleteById(id int) (bool, error)

	// 根據編號更新標籤
	UpdateById(id int, tag models.Tag) (bool, error)

	// 取得多個標籤
	GetByIds(id []string) ([]models.Tag, error)
}
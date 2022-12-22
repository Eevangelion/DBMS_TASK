package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type ITag interface {
	GetTagByID(TagID int) (tagOut *models.Tag, err error)
	Create(tag *models.Tag) (tagOut *models.Tag, err error)
	Delete(tag_id int) (err error)
}

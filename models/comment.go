package models

import (
	adapter "campground_go/adapters"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Comment        string `gorm:"not null"`
	UserID         uint   `gorm:"not null"`
	CampgroundName string `gorm:"not null"`
	Owner          User
}

type CommentDB struct {
	postgres *adapter.Postgres
}

func NewComment() *CommentDB {
	postgres := adapter.NewPostgresAdapter()
	// postgres.DestructiveReset(&Comment{})
	postgres.AutoMigrate(&Comment{})
	return &CommentDB{
		postgres: postgres,
	}
}

func (c *CommentDB) CreateComment(comment *Comment) error {
	return c.postgres.Create(comment).Error
}

func (c *CommentDB) FindByCampgroundName(campgroundName string) (*[]Comment, error) {
	var comments []Comment
	if err := c.postgres.FindByQuery("campground_name = ?", campgroundName, &comments); err != nil {
		return nil, err
	}

	return &comments, nil
}

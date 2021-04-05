package models

import (
	adapter "campground_go/adapters"
	"errors"

	"github.com/jinzhu/gorm"
)

type CampgroundDB struct {
	postgres *adapter.Postgres
}

type Campground struct {
	gorm.Model
	Name        string `gorm:"not null;unique_index"`
	Description string `gorm:"not null"`
	Image       string `gorm:"not null"`
}

func NewCampground() *CampgroundDB {
	postgres := adapter.NewPostgresAdapter()
	// postgres.DestructiveReset(&Campground{})
	return &CampgroundDB{
		postgres: postgres,
	}
}

func (c *CampgroundDB) Create(campground *Campground) error {
	if err := c.postgres.Create(campground).Error; err != nil {
		return err
	}
	return nil
}

func (c *CampgroundDB) FindByName(name string) (*Campground, error) {
	var campground Campground
	if err := c.postgres.FindByName(name, &campground); err != nil {
		return nil, errors.New("Could not find campground")
	}
	return &campground, nil
}

func (c *CampgroundDB) Find() (*[]Campground, error) {
	var campgrounds []Campground
	if err := c.postgres.Find(&campgrounds); err != nil {
		return nil, err
	}
	return &campgrounds, nil
}

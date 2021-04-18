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
	OwnerID     uint   `gorm:"not null"`
}

func NewCampground() *CampgroundDB {
	postgres := adapter.NewPostgresAdapter()
	// postgres.DestructiveReset(&Campground{})
	postgres.AutoMigrate(&Campground{})
	return &CampgroundDB{
		postgres: postgres,
	}
}

func (c *CampgroundDB) Create(campground *Campground) error {
	return c.postgres.Create(campground).Error
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

func (c *CampgroundDB) FindUsersCampgroundsByOwnerID(ownerID uint) (*[]Campground, error) {
	var campgrounds []Campground

	if err := c.postgres.FindByQuery("owner_id = ?", ownerID, &campgrounds); err != nil {
		return nil, err
	}

	return &campgrounds, nil
}

func (c *CampgroundDB) DeleteByName(name string) error {
	return c.postgres.Delete(name, &Campground{})
}

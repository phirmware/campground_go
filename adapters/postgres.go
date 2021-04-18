package adapter

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func connectDB(psqlInfo string) *gorm.DB {
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func NewPostgresAdapter() *Postgres {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db := connectDB(psqlInfo)
	db.LogMode(true)

	return &Postgres{
		DB: db,
	}
}

func (p *Postgres) AutoMigrate(model interface{}) error {
	if err := p.DB.AutoMigrate(model).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DestructiveReset(model interface{}) error {
	if err := p.DB.DropTableIfExists(model).Error; err != nil {
		return err
	}
	return p.AutoMigrate(model)
}

func (p *Postgres) Create(data interface{}) *gorm.DB {
	return p.DB.Create(data)
}

func (p *Postgres) FindByQuery(query string, queryValue interface{}, dst interface{}) error {
	return p.DB.Where(query, queryValue).Find(dst).Error
}

func (p *Postgres) FindByQueryAndPreload(query string, queryValue interface{}, dst interface{}, preload string) error {
	return p.DB.Preload(preload).Where(query, queryValue).Find(dst).Error
}

func (p *Postgres) FindByUsername(username string, dst interface{}) error {
	return p.DB.Where("username = ?", username).First(dst).Error
}

func (p *Postgres) FindByName(name string, dst interface{}) error {
	return p.DB.Where("name = ?", name).First(dst).Error
}

func (p *Postgres) Find(dst interface{}) error {
	return p.DB.Find(dst).Error
}

func (p *Postgres) Delete(name string, dst interface{}) error {
	return p.DB.Where("name = ?", name).Delete(dst).Error
}

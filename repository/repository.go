package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryI interface {
	Create(value any) error
	QueryAll(value any, where any) error
	Query(value any) error
	Update(value any) error
	Delete(value any) error
	Raw(result any, Query string, value ...any) error
}

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) Create(value any) error {
	return r.Db.Create(value).Error
}
func (r *Repository) QueryAll(value any, where any) error {
	return r.Db.Preload(clause.Associations).Where(where).Find(value).Error
}

func (r *Repository) Query(value any) error {
	return r.Db.Preload(clause.Associations).First(value).Error
}

func (r *Repository) Update(value any) error {
	return r.Db.Updates(value).Error
}

func (r *Repository) Delete(value any) error {
	return r.Db.Delete(value).Error
}

func (r *Repository) Raw(result any, Query string, value ...any) error {
	return r.Db.Raw(Query, value...).Scan(result).Error
}

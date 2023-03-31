package repository

import "gorm.io/gorm"

type RepositoryI interface {
	Create(value any) error
	QueryAll(value any) error
	Query(value any) error
	Update(value any) error
	Delete(value any) error
}

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) Create(value any) error {
	return r.Db.Create(value).Error
}
func (r *Repository) QueryAll(value any) error {
	return r.Db.Find(value).Error
}

func (r *Repository) Query(value any) error {
	return r.Db.First(value).Error
}

func (r *Repository) Update(value any) error {
	return r.Db.Updates(value).Error
}

func (r *Repository) Delete(value any) error {
	return r.Db.Delete(value).Error
}

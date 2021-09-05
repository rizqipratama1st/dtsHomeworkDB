package repository

import (
	"fmt"
	"homework/model"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func(r *MovieRepository) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	findResult := r.DB.Find(&movies)

	return movies, findResult.Error
}

func(r *MovieRepository) FindBySlug(slug string) (model.Movie, error) {
	var movies model.Movie
	findResult := r.DB.Where("slug = ?", slug).First(&movies)
	return movies, findResult.Error
}

func(r *MovieRepository) Save(movie model.Movie) (model.Movie, error) {
	trx := r.DB.Create(&movie)

	return movie, trx.Error
}

func(r *MovieRepository) Update(slug string, movie model.Movie) error {
	fmt.Println(movie)
	query := `UPDATE movies SET title = ?, description = ?, duration = ?, image = ?, slug = ? WHERE slug = ?`
	result := r.DB.Exec(query, movie.Title, movie.Description, movie.Duration, movie.Image, movie.Slug, slug)

	if result.Error != nil {
		return result.Error
	}

	// trx := r.DB.Where("slug = ?", slug).Update("title", "edit")
	// // trx := r.DB.Model(&model.Movie).Update(&movie)
	return nil
}

func(r *MovieRepository) DeleteBySlug(slug string) error {
	var movies model.Movie
	findResult := r.DB.Unscoped().Where("slug = ?", slug).Delete(&movies)
	if findResult.Error != nil {
		return findResult.Error
	}

	if findResult.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	
	return nil
}

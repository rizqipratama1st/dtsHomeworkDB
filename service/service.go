package service

import (
	"homework/model"
	"homework/repository"
)

type MovieService struct {
	MovieRepo repository.MovieRepository
}

func(s *MovieService) GetAllMovies() ([]model.Movie, error) {
	movies, err := s.MovieRepo.FindAll()

	return movies, err
}

func(s *MovieService) GetMoviesBySlug(slug string) (model.Movie, error) {
	movies, err := s.MovieRepo.FindBySlug(slug)

	return movies, err
}

func(s *MovieService) InsertMovie(movie model.Movie) (model.Movie,error) {

	response, err := s.MovieRepo.Save(movie)
	// fmt.Println(response)

	return response, err
}

func (s *MovieService) UpdateMovie(movies model.Movie, slug string) (model.Movie, error) {
	// update movie data
	var movie model.Movie
	err := s.MovieRepo.Update(slug, movies)
	if err != nil {
		return movie, err
	}

	movie, err = s.MovieRepo.FindBySlug(slug)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func(s *MovieService) DeleteMovieBySlug(slug string) error {
	err := s.MovieRepo.DeleteBySlug(slug)

	if err != nil {
		return err
	}

	return nil
}
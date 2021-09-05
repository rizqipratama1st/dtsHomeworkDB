package handler

import (
	"errors"
	"homework/model"
	"homework/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MovieHandler struct {
	MovieService service.MovieService
}

func(r *MovieHandler) GetMovies(c *fiber.Ctx) error {
	movies, err := r.MovieService.GetAllMovies()

	if err != nil {
		c.Status(403).JSON(err)
	}

	return c.JSON(movies)
}

func(r *MovieHandler) GetMoviesBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	response, err := r.MovieService.GetMoviesBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		// "msg":    "success update data",
		"result": response,
	})
	
}

func(r *MovieHandler) InsertMovie(c *fiber.Ctx) error {
	movie := model.Movie{}

	// if err := c.BodyParser(movie); err != nil {
	// 	c.Status(403).JSON(err)
	// }

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	response, err := r.MovieService.InsertMovie(movie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		// "msg":    "success update data",
		"result": response,
	})

}

// func (r *MovieHandler) UpdateMovie(c *fiber.Ctx) error {
//     fmt.Fprintf(c, "%s\n", c.Params("slug"))
//     return nil
// }

func (m *MovieHandler) UpdateMovie(c *fiber.Ctx) error {
	movie := model.Movie{}
	slug := c.Params("slug")
	// var mysqlErr *mysql.MySQLError

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	response, err := m.MovieService.UpdateMovie(movie, slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": response,
	})
}

func(r *MovieHandler) DeleteMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	err := r.MovieService.DeleteMovieBySlug(slug)
	if err != nil {
		var statusCode = fiber.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = fiber.StatusNotFound
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"result": "success delete data",
	})
	
}
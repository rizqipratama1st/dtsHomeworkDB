package main

import (
	"homework/handler"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"

	"homework/model"
	"homework/repository"
	"homework/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logdb "gorm.io/gorm/logger"
)

func helloMovie(c *fiber.Ctx) error {
	movie := model.Movie{
		Model: gorm.Model{},
		Title: "Titanic",
		Slug: "titanic",
		Description: "lorem",
		Duration: 324,
		Image: "knwekv",
	}
	return c.JSON(movie)
}

func helloPost(c *fiber.Ctx) error {
	movie := &model.Movie{}

	if err := c.BodyParser(movie); err != nil {
		c.Status(403).JSON(err)
	}
	return c.JSON(movie)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", helloMovie)
	app.Post("/", helloPost)

	log.SetLevel(log.DebugLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true) 
	
	gormDB, err := gorm.Open(mysql.Open("root:74712331@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{Logger: logdb.Default.LogMode(logdb.Info), 
	})

	if err != nil {
		panic(err.Error())
	}

	model.MigrateMovie(gormDB)

	movieRepo := repository.MovieRepository{
		DB: gormDB,
	}
	movieService := service.MovieService{
		MovieRepo: movieRepo,
	}
	movieHandler := handler.MovieHandler{
		MovieService: movieService,
	}

	app.Get("/movie", movieHandler.GetMovies)
	app.Get("/movie/:slug", movieHandler.GetMoviesBySlug)
	app.Post("/movie", movieHandler.InsertMovie)
	app.Put("/movie/:slug", movieHandler.UpdateMovie)
	app.Delete("/movie/:slug", movieHandler.DeleteMovieBySlug)

	app.Listen(":3000")
}

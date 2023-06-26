package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ishanshre/goFiberApiExample/internals/config"
	"github.com/ishanshre/goFiberApiExample/internals/handler"
)

func Router(global *config.AppConfig, app *fiber.App) {
	app.Use(cors.New())
	app.Get("/", handler.Repo.AllProducts)
	app.Post("/create", handler.Repo.InsertProduct)
	app.Get("/detail/:id", handler.Repo.GetProductByID)
	app.Put("/detail/:id/update", handler.Repo.UpdateProduct)
	app.Delete("/detail/:id/delete", handler.Repo.DeleteProduct)
}

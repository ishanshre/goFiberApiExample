package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/goFiberApiExample/internals/config"
	"github.com/ishanshre/goFiberApiExample/internals/driver"
	"github.com/ishanshre/goFiberApiExample/internals/models"
	"github.com/ishanshre/goFiberApiExample/internals/repository"
	"github.com/ishanshre/goFiberApiExample/internals/repository/dbrepo"
)

type Repository struct {
	Global *config.AppConfig
	DB     repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		Global: a,
		DB:     dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) AllProducts(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit")
	offset := ctx.QueryInt("offset")
	if limit == 0 {
		limit = 10
	}
	products, err := m.DB.GetAllProducts(limit, offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("error in fetching products: %s", err.Error()),
		})
	}
	return ctx.Status(200).JSON(products)
}

func (m *Repository) InsertProduct(ctx *fiber.Ctx) error {
	var p *models.Product
	if err := ctx.BodyParser(&p); err != nil {
		return err
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	if err := m.DB.InsertProduct(p); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Data inserted successfully",
	})
}

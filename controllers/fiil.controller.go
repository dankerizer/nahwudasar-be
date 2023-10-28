package controllers

import (
	"nahwudasar-be/controllers/presenter"
	"nahwudasar-be/helper"
	"nahwudasar-be/pkg/entities"
	"nahwudasar-be/pkg/fiil"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FiilController struct {
	service fiil.InterfaceFiil
}

func NewFiilHandler(conn *gorm.DB) *FiilController {
	service := fiil.NewService(conn)
	return &FiilController{
		service: service,
	}
}

func (p FiilController) GetAll(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page")
	if err != nil {
		log.Printf("Warning %v", err)
	}

	// row, err := c.ParamsInt("row")
	// if err != nil {
	// 	log.Printf("Warning %v", err)
	// }
	filter := c.Queries()
	results, err := p.service.GetAll(page, filter)

	if err != nil {
		log.Errorf("Error %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&helper.ResponseError{
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		})
	}

	if len(results) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&helper.ResponseError{
			StatusCode: fiber.StatusNotFound,
			Message:    fiber.ErrNotFound.Message,
		})
	}

	count, err := p.service.Count()
	if err != nil {
		log.Printf("Warning %v", err)
	}

	var meta *presenter.Meta
	if page > 0 {
		meta = &presenter.Meta{
			Rows:        20,
			CurrentPage: page,
			Totals:      count,
		}
	}

	resultData := presenter.Result[entities.Fiil]{
		Data: results,
		Meta: meta,
	}

	return c.Status(fiber.StatusOK).JSON(resultData)
}

func (p FiilController) GetOne(c *fiber.Ctx) error {
	result, err := p.service.GetOne(c.Params("id"))
	if err != nil {
		log.Errorf("Error %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&helper.ResponseError{
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		})
	}
	if len(result.Root) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&helper.ResponseError{
			StatusCode: fiber.StatusNotFound,
			Message:    fiber.ErrNotFound.Message,
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

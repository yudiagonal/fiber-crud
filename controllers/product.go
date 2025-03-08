package controllers

import (
	"fiber-crud/interfaces"
	"fiber-crud/models"
	"fiber-crud/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	repo interfaces.ProductRepository
}

func NewProductController(repo interfaces.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	// ambil parameter pagination dari query string
	page, _ := strconv.Atoi(ctx.Query("page", "1"))   // default 1
	limit, _ := strconv.Atoi(ctx.Query("page", "10")) // default 10

	//ambil data dari repository
	products, total, err := c.repo.GetAllProducts(page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch product",
		})
	}

	// meta data pagination
	meta := utils.NewPaginationMeta(page, limit, total)
	// Kembalikan respons dengan metadata pagination
	return ctx.JSON(fiber.Map{
		"data": products,
		"meta": meta,
	})
}

func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	product := c.repo.GetProductByID(id)
	if product == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdProduct := c.repo.CreateProduct(product)
	return ctx.Status(fiber.StatusCreated).JSON(createdProduct)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var updatedProduct models.Product
	if err := ctx.BodyParser(&updatedProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	product := c.repo.UpdateProduct(id, updatedProduct)
	if product == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return ctx.JSON(product)
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if !c.repo.DeleteProduct(id) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

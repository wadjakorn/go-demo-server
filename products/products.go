package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductResponse struct {
	Title string `json:"title"`
}

func GetProductHandler(service *ProductService) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := service.GetAll()
		return c.JSON(http.StatusOK, r)
	}
}

type IRepository interface {
	GetProducts() ([]ProductResponse, error)
}

type ProductRepository struct {
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (r *ProductRepository) GetProducts() ([]ProductResponse, error) {
	var products []ProductResponse
	products = append(products, ProductResponse{Title: "Nuphy Air75"}, ProductResponse{Title: "Loga Kirin Pro"})
	return products, nil
}

type ProductService struct {
	repo IRepository
}

func NewProductService(repo IRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) GetAll() []ProductResponse {
	result, _ := p.repo.GetProducts()
	return result
}
package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductContoller(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 300,
		},
	}

	ctx.JSON(http.StatusOK, products)
}

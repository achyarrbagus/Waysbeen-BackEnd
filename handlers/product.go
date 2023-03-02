package handlers

import (
	productsdto "backEnd/dto/products"
	dto "backEnd/dto/result"
	"backEnd/models"
	"strconv"

	"backEnd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandleProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertProduct(product)})
}

func (h *handlerProduct) FindProduct(c echo.Context) error {
	product, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	request := new(productsdto.CreateProductRequset)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Stock:       request.Stock,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertProduct(data)})

}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data, err := h.ProductRepository.DeleteProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertProduct(data)})
}

func (h *handlerProduct) UpdateProduct(C echo.Context) error {
	request := new(productsdto.UpdateProductRequest)
	if err := C.Bind(&request); err != nil {
		return C.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(C.Param("id"))
	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return C.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return C.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return C.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertProduct(data)})

}

func convertProduct(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
		Stock:       u.Stock,
	}
}

package handlers

import (
	productsdto "backEnd/dto/products"
	dto "backEnd/dto/result"
	"backEnd/models"
	"fmt"
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

	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	price, _ := strconv.Atoi(c.FormValue("price"))
	qty, _ := strconv.Atoi(c.FormValue("stock"))
	// category_id, _ := strconv.Atoi(c.FormValue("category_id"))

	product := models.Product{
		Name:        c.FormValue("name"),
		Price:       price,
		Description: c.FormValue("desc"),
		Stock:       qty,
		Photo:       dataFile,
		UserID:      1,
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

// import (
// 	"backEnd/dto"
// 	"backEnd/dto/result"
// 	"backEnd/models"
// 	"backEnd/repositories"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-playground/validator"
// 	"github.com/labstack/echo/v4"
// )

// type productControl struct {
// 	ProductRepository repositories.ProductRepository
// }

// func ControlProduct(ProductRepository repositories.ProductRepository) *productControl {
// 	return &productControl{ProductRepository}
// }

// func (h *productControl) FindProducts(c echo.Context) error {
// 	products, err := h.ProductRepository.FindProduct()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: products})
// }

// func (h *productControl) GetProducts(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	products, err := h.ProductRepository.GetProducts(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(products)})
// }

// func (h *productControl) CreateProduct(c echo.Context) error {
// 	// get the datafile here
// 	dataFile := c.Get("dataFile").(string)
// 	fmt.Println("this is data file", dataFile)

// 	price, _ := strconv.Atoi(c.FormValue("price"))
// 	stock, _ := strconv.Atoi(c.FormValue("stock"))

// 	request := dto.CreateProductRequest{
// 		Name:        c.FormValue("name"),
// 		Description: c.FormValue("desc"),
// 		Price:       price,
// 		Photo:       dataFile,
// 		Stock:       stock,
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	// data form pattern submit to pattern entity db product
// 	product := models.Product{
// 		Name:        request.Name,
// 		Price:       request.Price,
// 		Description: request.Description,
// 		Stock:       request.Stock,
// 		Photo:       request.Photo,
// 	}

// 	data, err := h.ProductRepository.CreateProduct(product)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
// }

// func (h *productControl) UpdateProduct(c echo.Context) error {
// 	request := new(dto.UpdateProductRequest)
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	product, err := h.ProductRepository.GetProducts(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	if request.Name != "" {
// 		product.Name = request.Name
// 	}
// 	if request.Price != 0 {
// 		product.Price = request.Price
// 	}
// 	if request.Description != "" {
// 		product.Description = request.Description
// 	}
// 	if request.Stock != 0 {
// 		product.Stock = request.Stock
// 	}
// 	if request.Photo != "" {
// 		product.Photo = request.Photo
// 	}

// 	data, err := h.ProductRepository.UpdateProduct(product, id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
// }

// func (h *productControl) DeleteProduct(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user, err := h.ProductRepository.GetProducts(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	data, err := h.ProductRepository.DeleteProduct(user, id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
// }

// func convProduct(u models.Product) models.ProductResponse {
// 	return models.ProductResponse{
// 		Name:        u.Name,
// 		Price:       u.Price,
// 		Description: u.Description,
// 		Stock:       u.Stock,
// 		Photo:       u.Photo,
// 		User:        u.User,
// 	}
// }

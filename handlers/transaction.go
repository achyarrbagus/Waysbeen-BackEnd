package handlers

import (
	dto "backEnd/dto/result"
	transactiondto "backEnd/dto/transaction"
	"backEnd/models"
	"backEnd/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
	CartRepository        repositories.CartRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository, CartRepository repositories.CartRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository, CartRepository}
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["id"].(float64)
	// data form pattern submit to pattern entity db transaction
	transaction := models.Transaction{
		UserID:    int(isLoginUser),
		Name:      request.Name,
		Email:     request.Email,
		Status:    "Waiting Approve",
		Address:   request.Address,
		Date:      time.Time{},
		CreatedAt: time.Time{},
	}

	newTransaction, err := h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	fmt.Println(request.Product)
	for _, item := range request.Product {
		newCart := models.Cart{
			OrderQuantity: item.OrderQuantity,
			TransactionID: newTransaction.ID,
			ProductID:     item.ID,
		}
		fmt.Println(newCart)
		newCart, err := h.CartRepository.CreateCart(newCart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: newTransaction})

}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

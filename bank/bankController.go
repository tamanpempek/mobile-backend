package bank

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	bankService BankService
}

func NewController(bankService BankService) *controller {
	return &controller{bankService}
}

func (cn *controller) GetBanks(c *gin.Context) {
	banks, err := cn.bankService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var banksResponse []BankResponse

	for _, bank := range banks {
		bankResponse := convertToBankResponse(bank)

		banksResponse = append(banksResponse, bankResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  banksResponse,
	})
}

func (cn *controller) GetAdminBanks(c *gin.Context) {
	banks, err := cn.bankService.FindAdminBanks()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var banksResponse []BankResponse

	for _, bank := range banks {
		bankResponse := convertToBankResponse(bank)

		banksResponse = append(banksResponse, bankResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  banksResponse,
	})
}

func (cn *controller) GetBanksByUser(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid user ID",
		})
		return
	}

	banks, err := cn.bankService.FindBanksByUser(userId)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Banks not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	var banksResponse []BankResponse

	for _, bank := range banks {
		bankResponse := convertToBankResponse(bank)

		banksResponse = append(banksResponse, bankResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  banksResponse,
	})
}

func (cn *controller) GetBank(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid bank ID",
		})
		return
	}

	bank, err := cn.bankService.FindBankByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Bank not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToBankResponse(bank),
	})
}

func (cn *controller) CreateBank(c *gin.Context) {
	var bankRequest BankCreateRequest

	err := c.ShouldBindJSON(&bankRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	bank, err := cn.bankService.CreateBank(bankRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToBankResponse(bank),
	})
}

func (cn *controller) UpdateBank(c *gin.Context) {
	var bankRequest BankUpdateRequest

	err := c.ShouldBindJSON(&bankRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid bank ID",
		})
		return
	}

	bank, err := cn.bankService.UpdateBank(id, bankRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Bank not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToBankResponse(bank),
	})
}

func (ch *controller) DeleteBank(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid bank ID",
		})
		return
	}

	bank, err := ch.bankService.DeleteBank(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Bank not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToBankResponse(bank),
	})
}

func convertToBankResponse(bank Bank) BankResponse {
	return BankResponse{
		ID:     bank.ID,
		UserID: bank.UserID,
		Type:   bank.Type,
		Name:   bank.Name,
		Number: bank.Number,
	}
}

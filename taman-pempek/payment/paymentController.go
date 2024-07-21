package payment

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type controller struct {
	paymentService PaymentService
}

func NewController(paymentService PaymentService) *controller {
	return &controller{paymentService}
}

func (cn *controller) GetPayments(c *gin.Context) {
	payments, err := cn.paymentService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var paymentsResponse []PaymentResponse

	for _, payment := range payments {
		paymentResponse := convertToPaymentResponse(payment)

		paymentsResponse = append(paymentsResponse, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  paymentsResponse,
	})
}

func (cn *controller) GetPayment(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := cn.paymentService.FindPaymentByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func (cn *controller) GetPaymentByUserAndStatus(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid payment ID",
		})
		return
	}

	paymentStatus := c.Param("paymentStatus")

	payments, err := cn.paymentService.FindPaymentByUserAndStatus(userId, paymentStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var paymentsResponse []PaymentResponse

	for _, payment := range payments {
		paymentResponse := convertToPaymentResponse(payment)

		paymentsResponse = append(paymentsResponse, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  paymentsResponse,
	})
}

func (cn *controller) GetPaymentByStatus(c *gin.Context) {
	paymentStatus := c.Param("paymentStatus")

	payments, err := cn.paymentService.FindPaymentByStatus(paymentStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var paymentsResponse []PaymentResponse

	for _, payment := range payments {
		paymentResponse := convertToPaymentResponse(payment)

		paymentsResponse = append(paymentsResponse, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  paymentsResponse,
	})
}

func (cn *controller) CreatePayment(c *gin.Context) {
	var paymentRequest PaymentCreateRequest

	apiKey := goDotEnvVariable("APIKEY")
	apiSecret := goDotEnvVariable("APISECRET")

	urlCloudinary := "cloudinary://" + apiKey + ":" + apiSecret + "@dqudegiey"

	err := c.ShouldBind(&paymentRequest)

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

	file, err := paymentRequest.Image.Open()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	ctx := context.Background()

	cldService, err := cloudinary.NewFromURL(urlCloudinary)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	imageResponse, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	paymentRequest.Image.Filename = imageResponse.SecureURL

	payment, err := cn.paymentService.CreatePayment(paymentRequest)

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
		"data":  convertToPaymentResponse(payment),
	})
}

func (cn *controller) UpdatePayment(c *gin.Context) {
	var paymentRequest PaymentUpdateRequest

	err := c.ShouldBind(&paymentRequest)

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
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := cn.paymentService.UpdatePayment(id, paymentRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func (ch *controller) DeletePayment(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := ch.paymentService.DeletePayment(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func convertToPaymentResponse(payment Payment) PaymentResponse {
	return PaymentResponse{
		ID:            payment.ID,
		UserID:        payment.UserID,
		DeliveryID:    payment.DeliveryID,
		TotalPrice:    payment.TotalPrice,
		Image:         payment.Image,
		Address:       payment.Address,
		Whatsapp:      payment.Whatsapp,
		PaymentStatus: payment.PaymentStatus,
		DeliveryName:  payment.DeliveryName,
		Resi:          payment.Resi,
	}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

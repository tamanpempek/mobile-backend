package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	deliveryService DeliveryService
}

func NewController(deliveryService DeliveryService) *controller {
	return &controller{deliveryService}
}

func (cn *controller) GetDeliveries(c *gin.Context) {
	deliveries, err := cn.deliveryService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var deliveriesResponse []DeliveryResponse

	for _, delivery := range deliveries {
		deliveryResponse := convertToDeliveryResponse(delivery)

		deliveriesResponse = append(deliveriesResponse, deliveryResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  deliveriesResponse,
	})
}

func (cn *controller) GetDelivery(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid delivery ID",
		})
		return
	}

	delivery, err := cn.deliveryService.FindDeliveryByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "delivery not found" {
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
		"data":  convertToDeliveryResponse(delivery),
	})
}

func (cn *controller) CreateDelivery(c *gin.Context) {
	var deliveryRequest DeliveryCreateRequest

	err := c.ShouldBindJSON(&deliveryRequest)

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

	delivery, err := cn.deliveryService.CreateDelivery(deliveryRequest)

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
		"data":  convertToDeliveryResponse(delivery),
	})
}

func (cn *controller) UpdateDelivery(c *gin.Context) {
	var deliveryRequest DeliveryUpdateRequest

	err := c.ShouldBindJSON(&deliveryRequest)

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
			"msg":   "Invalid delivery ID",
		})
		return
	}

	delivery, err := cn.deliveryService.UpdateDelivery(id, deliveryRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Delivery not found" {
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
		"data":  convertToDeliveryResponse(delivery),
	})
}

func (ch *controller) DeleteDelivery(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid delivery ID",
		})
		return
	}

	delivery, err := ch.deliveryService.DeleteDelivery(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Delivery not found" {
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
		"data":  convertToDeliveryResponse(delivery),
	})
}

func convertToDeliveryResponse(delivery Delivery) DeliveryResponse {
	return DeliveryResponse{
		ID:       delivery.ID,
		Name:     delivery.Name,
		Whatsapp: delivery.Whatsapp,
	}
}

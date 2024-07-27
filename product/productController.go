package product

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
	productService ProductService
}

func NewController(productService ProductService) *controller {
	return &controller{productService}
}

func (cn *controller) GetProducts(c *gin.Context) {
	products, err := cn.productService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productResponse := convertToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  productsResponse,
	})
}

func (cn *controller) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid product ID",
		})
		return
	}

	product, err := cn.productService.FindProductByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Product not found" {
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
		"data":  convertToProductResponse(product),
	})
}

func (cn *controller) GetProductByUserIDAndCategoryID(c *gin.Context) {
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

	categoryIdString := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid category ID",
		})
		return
	}

	products, err := cn.productService.GetProductByUserIDAndCategoryID(userId, categoryId)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Products not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productResponse := convertToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  productsResponse,
	})
}

func (cn *controller) GetProductByUser(c *gin.Context) {
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

	products, err := cn.productService.GetProductByUser(userId)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Products not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productResponse := convertToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  productsResponse,
	})
}

func (cn *controller) GetProductByCategory(c *gin.Context) {
	categoryIdString := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid user ID",
		})
		return
	}

	products, err := cn.productService.GetProductByCategory(categoryId)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Products not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productResponse := convertToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  productsResponse,
	})
}

func (cn *controller) CreateProduct(c *gin.Context) {
	var productRequest ProductCreateRequest

	apiKey := goDotEnvVariable("APIKEY")
	apiSecret := goDotEnvVariable("APISECRET")

	urlCloudinary := "cloudinary://" + apiKey + ":" + apiSecret + "@dqudegiey"

	err := c.ShouldBind(&productRequest)

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

	file, err := productRequest.Image.Open()

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

	productRequest.Image.Filename = imageResponse.SecureURL

	product, err := cn.productService.CreateProduct(productRequest)

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
		"data":  convertToProductResponse(product),
	})
}

func (cn *controller) UpdateProduct(c *gin.Context) {
	var productRequest ProductUpdateRequest

	apiKey := goDotEnvVariable("APIKEY")
	apiSecret := goDotEnvVariable("APISECRET")

	urlCloudinary := "cloudinary://" + apiKey + ":" + apiSecret + "@dqudegiey"

	err := c.ShouldBind(&productRequest)

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

	if productRequest.Image != nil {
		file, _ := productRequest.Image.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(urlCloudinary)
		imageResponse, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		productRequest.Image.Filename = imageResponse.SecureURL
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid product ID",
		})
		return
	}

	product, err := cn.productService.UpdateProduct(id, productRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Product not found" {
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
		"data":  convertToProductResponse(product),
	})
}

func (ch *controller) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid product ID",
		})
		return
	}

	product, err := ch.productService.DeleteProduct(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Product not found" {
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
		"data":  convertToProductResponse(product),
	})
}

func convertToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

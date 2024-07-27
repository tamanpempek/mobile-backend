package category

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	categoryService CategoryService
}

func NewController(categoryService CategoryService) *controller {
	return &controller{categoryService}
}

func (cn *controller) GetCategories(c *gin.Context) {
	categories, err := cn.categoryService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var categoriesResponse []CategoryResponse

	for _, category := range categories {
		categoryResponse := convertToCategoryResponse(category)

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  categoriesResponse,
	})
}

func (cn *controller) GetCategory(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid category ID",
		})
		return
	}

	category, err := cn.categoryService.FindCategoryByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Category not found" {
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
		"data":  convertToCategoryResponse(category),
	})
}

func (cn *controller) CreateCategory(c *gin.Context) {
	var categoryRequest CategoryCreateRequest

	err := c.ShouldBindJSON(&categoryRequest)

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

	category, err := cn.categoryService.CreateCategory(categoryRequest)

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
		"data":  convertToCategoryResponse(category),
	})
}

func (cn *controller) UpdateCategory(c *gin.Context) {
	var categoryRequest CategoryUpdateRequest

	err := c.ShouldBindJSON(&categoryRequest)

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
			"msg":   "Invalid category ID",
		})
		return
	}

	category, err := cn.categoryService.UpdateCategory(id, categoryRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Category not found" {
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
		"data":  convertToCategoryResponse(category),
	})
}

func (ch *controller) DeleteCategory(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid category ID",
		})
		return
	}

	category, err := ch.categoryService.DeleteCategory(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Category not found" {
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
		"data":  convertToCategoryResponse(category),
	})
}

func convertToCategoryResponse(category Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

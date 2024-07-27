package setting

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
	settingService SettingService
}

func NewController(settingService SettingService) *controller {
	return &controller{settingService}
}

func (cn *controller) GetSetting(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid setting ID",
		})
		return
	}

	setting, err := cn.settingService.FindSettingByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Setting not found" {
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
		"data":  convertToSettingResponse(setting),
	})
}

func (cn *controller) UpdateSetting(c *gin.Context) {
	var settingRequest SettingUpdateRequest

	apiKey := goDotEnvVariable("APIKEY")
	apiSecret := goDotEnvVariable("APISECRET")

	urlCloudinary := "cloudinary://" + apiKey + ":" + apiSecret + "@dqudegiey"

	err := c.ShouldBind(&settingRequest)

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

	if settingRequest.Image != nil {
		file, _ := settingRequest.Image.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(urlCloudinary)
		imageResponse, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		settingRequest.Image.Filename = imageResponse.SecureURL
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid setting ID",
		})
		return
	}

	setting, err := cn.settingService.UpdateSetting(id, settingRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Setting not found" {
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
		"data":  convertToSettingResponse(setting),
	})
}

func convertToSettingResponse(setting Setting) SettingResponse {
	return SettingResponse{
		ID:          setting.ID,
		Image:       setting.Image,
		Description: setting.Description,
		Email:       setting.Email,
		Instagram:   setting.Instagram,
		Website:     setting.Website,
	}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

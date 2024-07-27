package main

import (
	"log"
	"os"
	"taman-pempek/bank"
	"taman-pempek/cart"
	"taman-pempek/category"
	"taman-pempek/delivery"
	"taman-pempek/middleware"
	"taman-pempek/payment"
	"taman-pempek/product"
	"taman-pempek/setting"
	"taman-pempek/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	dbUser := goDotEnvVariable("MYSQLUSER")
	dbPassword := goDotEnvVariable("MYSQLPASSWORD")
	dbHost := goDotEnvVariable("MYSQLHOST")
	dbPort := goDotEnvVariable("MYSQLPORT")
	dbName := goDotEnvVariable("MYSQLDATABASE")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	migration(db)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userMiddleware := middleware.NewMiddleware(userService)

	requireAuth := userMiddleware.RequireAuth

	routeUser(db, v1, requireAuth)
	routeProduct(db, v1, requireAuth)
	routeBank(db, v1, requireAuth)
	routeCategory(db, v1, requireAuth)
	routeDelivery(db, v1, requireAuth)
	routeCart(db, v1, requireAuth)
	routePayment(db, v1, requireAuth)
	routeSetting(db, v1, requireAuth)

	router.Run(":8888") // port
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&cart.Cart{})
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&delivery.Delivery{})
	db.AutoMigrate(&payment.Payment{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&setting.Setting{})
}

func routeUser(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	v.GET("/users", userController.GetUsers)
	v.GET("/users/role/:role", userController.FindUsersByRole)
	v.GET("/user/:id", userController.GetUser)
	v.POST("/user/register", userController.CreateUser)
	v.PUT("/user/update/:id", userController.UpdateUser)
	v.DELETE("/user/delete/:id", userController.DeleteUser)

	v.POST("/login", userController.Login)
	v.POST("/logout", userController.Logout)
}

func routeProduct(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	v.GET("/products", productController.GetProducts)
	v.GET("/products/:userId/", productController.GetProductByUser)
	v.GET("/products/category/:categoryId", productController.GetProductByCategory)
	v.GET("/products/:userId/:categoryId", productController.GetProductByUserIDAndCategoryID)
	v.GET("/product/:id", productController.GetProduct)
	v.POST("/product/create", productController.CreateProduct)
	v.PUT("/product/update/:id", productController.UpdateProduct)
	v.DELETE("/product/delete/:id", productController.DeleteProduct)
}

func routeBank(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	bankRepository := bank.NewRepository(db)
	bankService := bank.NewService(bankRepository)
	bankController := bank.NewController(bankService)

	v.GET("/banks", bankController.GetBanks)
	v.GET("/banks/admin", bankController.GetAdminBanks)
	v.GET("/banks/:userId", bankController.GetBanksByUser)
	v.GET("/bank/:id", bankController.GetBank)
	v.POST("/bank/create", bankController.CreateBank)
	v.PUT("/bank/update/:id", bankController.UpdateBank)
	v.DELETE("/bank/delete/:id", bankController.DeleteBank)
}

func routeCategory(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryController := category.NewController(categoryService)

	v.GET("/categories", categoryController.GetCategories)
	v.GET("/category/:id", categoryController.GetCategory)
	v.POST("/category/create", categoryController.CreateCategory)
	v.PUT("/category/update/:id", categoryController.UpdateCategory)
	v.DELETE("/category/delete/:id", categoryController.DeleteCategory)
}

func routeDelivery(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	deliveryRepository := delivery.NewRepository(db)
	deliveryService := delivery.NewService(deliveryRepository)
	deliveryController := delivery.NewController(deliveryService)

	v.GET("/deliveries", deliveryController.GetDeliveries)
	v.GET("/delivery/:id", deliveryController.GetDelivery)
	v.POST("/delivery/create", deliveryController.CreateDelivery)
	v.PUT("/delivery/update/:id", deliveryController.UpdateDelivery)
	v.DELETE("/delivery/delete/:id", deliveryController.DeleteDelivery)
}

func routeCart(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	cartRepository := cart.NewRepository(db)
	cartService := cart.NewService(cartRepository)
	cartController := cart.NewController(cartService)

	v.GET("/carts", cartController.GetCarts)
	v.GET("/cart/:id", cartController.GetCart)
	v.GET("/carts/payment/:paymentId", cartController.FindCartsByPaymentID)
	v.GET("/carts/product/:productId", cartController.FindCartsByProductID)
	v.GET("/carts/:isActived/:userId", cartController.FindStatusCardByUser)
	v.GET("/carts/total/:isActived/:userId", cartController.SumTotalPriceByUser)
	v.POST("/cart/create", cartController.CreateCart)
	v.PUT("/cart/update/:id", cartController.UpdateCart)
	v.DELETE("/cart/delete/:id", cartController.DeleteCart)
}

func routePayment(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentController := payment.NewController(paymentService)

	v.GET("/payments", paymentController.GetPayments)
	v.GET("/payments/:userId/:paymentStatus", paymentController.GetPaymentByUserAndStatus)
	v.GET("/payments/status/:paymentStatus", paymentController.GetPaymentByStatus)
	v.GET("/payment/:id", paymentController.GetPayment)
	v.POST("/payment/create", paymentController.CreatePayment)
	v.PUT("/payment/update/:id", paymentController.UpdatePayment)
	v.DELETE("/payment/delete/:id", paymentController.DeletePayment)
}

func routeSetting(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	settingRepository := setting.NewRepository(db)
	settingService := setting.NewService(settingRepository)
	settingController := setting.NewController(settingService)

	v.GET("/setting/:id", settingController.GetSetting)
	v.PUT("/setting/update/:id", settingController.UpdateSetting)
}

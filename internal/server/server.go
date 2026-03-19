package server

import (
	"net/http"

	"github.com/Adwait-aayush/ECMCP/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	logger zerolog.Logger
}

func New(cfg *config.Config, db *gorm.DB, logger zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) SetupRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	router.GET("/health", s.healthchecker)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.register)
			auth.POST("/login", s.login)
			auth.POST("/refresh", s.refreshToken)
			auth.POST("/logout", s.logout)
		}
		api.GET("/categories", s.getCategories)
		api.GET("/products", s.getProducts)
		api.GET("/products/:id", s.getProduct)

		protected := api.Group("/")
		protected.Use(s.authMiddleware())
		{
			user := protected.Group("/users")
			{
				user.GET("/profile", s.getProfile)
				user.PUT("/profile", s.updateProfile)
			}
			category := protected.Group("/categories")
			{
				category.POST("/", s.adminMiddleware(), s.createCategory)
				category.PUT("/:id", s.adminMiddleware(), s.updateCategory)
				category.DELETE("/:id", s.adminMiddleware(), s.deleteCategory)
			}
			product := protected.Group("/products")
			{
				product.POST("/", s.adminMiddleware(), s.createProduct)
				product.PUT("/:id", s.adminMiddleware(), s.updateProduct)
				product.DELETE("/:id", s.adminMiddleware(), s.deleteProduct)
			}

		}
	}

	return router

}

func (s *Server) healthchecker(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}

}

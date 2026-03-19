package server

import (
	"strconv"

	"github.com/Adwait-aayush/ECMCP/internal/dto"
	"github.com/Adwait-aayush/ECMCP/internal/services"
	"github.com/Adwait-aayush/ECMCP/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	productService := services.NewProductService(s.db)
	resp, err := productService.CreateCategory(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to create category", err)
		return
	}
	utils.CreatedResponse(c, "Category Created Successfully", resp)
}

func (s *Server) getCategories(c *gin.Context) {
	productService := services.NewProductService(s.db)
	resp, err := productService.GetCategories()
	if err != nil {
		utils.BadRequestResponse(c, "Failed to fetch categories", err)
		return
	}
	utils.SuccessResponse(c, "Categories fetched successfully", resp)
}

func (s *Server) updateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid category id", err)
		return
	}
	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid Request Data", err)
		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update Category", err)
		return
	}
	utils.SuccessResponse(c, "Category Created Successfully", category)

}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteCategory(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete category", err)
		return
	}

	utils.SuccessResponse(c, "CategoryDeleted Successfully", nil)
}

func (s *Server) createProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid Request Data", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create product", err)
		return
	}
	utils.SuccessResponse(c, "Product Created Successfully", product)

}

func (s *Server) getProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		utils.BadRequestResponse(c, "invalid page", err)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		utils.BadRequestResponse(c, "invalid limit", err)
		return
	}

	productservice := services.NewProductService(s.db)
	products, meta, err := productservice.GetProducts(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "FAILED TO FETCH", err)
		return
	}
	utils.PaginatedSuccessResponse(c, "Products retrieved Successfully", products, *meta)

}

func (s *Server) getProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product id", err)
		return
	}
	productServices := services.NewProductService(s.db)
	product, err := productServices.GetProduct(uint(id))

	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to load product", err)
		return
	}
	utils.SuccessResponse(c, "loaded product successfully", product)
}

func (s *Server) updateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product id", err)
		return
	}

	var req dto.UpdateProductRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid product data", err)
		return
	}
	productService := services.NewProductService(s.db)

	resp, err := productService.UpdateProduct(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Unable to Update product", err)
		return
	}
	utils.SuccessResponse(c, "updated product Successfully", resp)
}

func (s *Server) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product id", err)
		return
	}
	productService := services.NewProductService(s.db)
	err = productService.DeleteProduct(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete product", err)
		return
	}
	utils.SuccessResponse(c, "Deleted product Successfully", nil)
}

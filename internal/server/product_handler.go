package server

import (
	"strconv"

	"github.com/Adwait-aayush/ECMCP/internal/dto"
	"github.com/Adwait-aayush/ECMCP/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	resp, err := s.productService.CreateCategory(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to create category", err)
		return
	}
	utils.CreatedResponse(c, "Category Created Successfully", resp)
}

func (s *Server) getCategories(c *gin.Context) {

	resp, err := s.productService.GetCategories()
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

	category, err := s.productService.UpdateCategory(uint(id), &req)
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

	if err := s.productService.DeleteCategory(uint(id)); err != nil {
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

	product, err := s.productService.CreateProduct(&req)
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

	products, meta, err := s.productService.GetProducts(page, limit)
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

	product, err := s.productService.GetProduct(uint(id))

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

	resp, err := s.productService.UpdateProduct(uint(id), &req)
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

	err = s.productService.DeleteProduct(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete product", err)
		return
	}
	utils.SuccessResponse(c, "Deleted product Successfully", nil)
}

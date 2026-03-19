package server

import (
	"github.com/Adwait-aayush/ECMCP/internal/dto"
	"github.com/Adwait-aayush/ECMCP/internal/services"
	"github.com/Adwait-aayush/ECMCP/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	authService := services.NewAuthService(s.db, s.config)
	resp, err := authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to register user", err)
		return
	}
	utils.CreatedResponse(c, "User Registered Successfully", resp)
}

func (s *Server) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	authService := services.NewAuthService(s.db, s.config)
	resp, err := authService.Login(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to login", err)
		return
	}
	utils.SuccessResponse(c, "Login Successful", resp)
}

func (s *Server) refreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	authService := services.NewAuthService(s.db, s.config)
	resp, err := authService.RefreshToken(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to refresh token", err)
		return
	}
	utils.SuccessResponse(c, "Token Refreshed Successfully", resp)
}

func (s *Server) logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	authService := services.NewAuthService(s.db, s.config)
	if err := authService.Logout(req.RefreshToken); err != nil {
		utils.BadRequestResponse(c, "Failed to logout", err)
		return
	}
	utils.SuccessResponse(c, "Logout Successful", nil)
}

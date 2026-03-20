package server

import (
	"github.com/Adwait-aayush/ECMCP/internal/dto"
	"github.com/Adwait-aayush/ECMCP/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	resp, err := s.authService.Register(&req)
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
	resp, err := s.authService.Login(&req)
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

	resp, err := s.authService.RefreshToken(&req)
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

	if err := s.authService.Logout(req.RefreshToken); err != nil {
		utils.BadRequestResponse(c, "Failed to logout", err)
		return
	}
	utils.SuccessResponse(c, "Logout Successful", nil)
}

func (s *Server) getProfile(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		utils.UnauthorizedResponse(c, "User ID not found", nil)
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		utils.BadRequestResponse(c, "Invalid user ID type", nil)
		return
	}

	profile, err := s.userService.GetProfile(userID)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to get profile", err)
		return
	}
	utils.SuccessResponse(c, "Profile retrieved successfully", profile)
}

func (s *Server) updateProfile(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		utils.UnauthorizedResponse(c, "User ID not found", nil)
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		utils.BadRequestResponse(c, "Invalid user ID type", nil)
		return
	}
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	profile, err := s.userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to update profile", err)
		return
	}
	utils.SuccessResponse(c, "Profile updated successfully", profile)
}

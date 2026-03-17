package server

import (
	"strings"

	"github.com/Adwait-aayush/ECMCP/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthHeader := c.GetHeader("Authorization")
		if AuthHeader == "" {
			utils.UnauthorizedResponse(c, "Authorization header missing", nil)
			c.Abort()
			return
		}

		tokenparts := strings.SplitN(AuthHeader, " ", 2)
		if len(tokenparts) != 2 || tokenparts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Invalid authorization header format", nil)

		}
		claims,err:=utils.ValidateToken(tokenparts[1],s.config.JWT.Secret)
		if err!=nil{
			utils.UnauthorizedResponse(c, "Invalid or expired token", nil)
			c.Abort()
			return
		}
		c.Set("user_id",claims.UserID)
		c.Set("user_email",claims.Email)
		c.Set("user_role",claims.Role)
		c.Next()
	}
}

func (s *Server) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != "admin" {
			utils.ForbiddenResponse(c, "Admin access required", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

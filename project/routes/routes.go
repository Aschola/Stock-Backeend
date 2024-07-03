package routes

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
	"project/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes (no authentication required)
	r.POST("/admin/login", controllers.AdminLogin) // Admin login endpoint

	// Super admin routes (no authentication required for signup)
	superadmin := r.Group("/superadmin")
	{
		superadmin.POST("/signup", controllers.SuperAdminSignup) // Super admin signup endpoint
		superadmin.Use(middlewares.AuthMiddleware())             // Apply auth middleware for subsequent endpoints
		superadmin.POST("/addadmin", controllers.AddAdmin)       // Add admin endpoint (requires authentication)
	}

	// Authenticated routes (for admins and super admins)
	auth := r.Group("/admin")
	auth.Use(middlewares.AuthMiddleware()) // Apply authentication middleware for admin routes
	{
		auth.POST("/users", controllers.AddUser)       // Add user endpoint
		auth.PUT("/users/:id", controllers.EditUser)   // Edit user endpoint
		auth.GET("/users/:id", controllers.GetUserByID) // Get user by ID endpoint
		auth.DELETE("/users/:id", controllers.DeleteUser) // Delete user endpoint
	}

	return r
}

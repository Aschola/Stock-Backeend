package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project/models"
	"project/utils"
)

// SuperAdminSignup allows the super admin to sign up
func SuperAdminSignup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure only super admin can sign up
	if input.RoleID != 1 { // Assuming RoleID 1 is for superadmin
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only super admin can sign up"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = hashedPassword

	// Save user to database
	if err := utils.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(input.ID, input.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respond with token
	c.JSON(http.StatusOK, gin.H{"token": token})
}


// AddAdmin allows the super admin to add an admin to the system
func AddAdmin(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure only super admin can add admin
	userID, _ := c.Get("userID")
	superAdminID := userID.(uint) // Assuming userID is set in context by auth middleware

	// Retrieve super admin from database (optional: you may need to check if user exists and is a super admin)
	var superAdmin models.User
	if err := utils.DB.First(&superAdmin, superAdminID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Super admin not found"})
		return
	}

	// Check if super admin is authorized
	if superAdmin.RoleID != 1 { // Assuming RoleID 1 is for superadmin
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only super admin can add admin"})
		return
	}

	// Set role for new admin
	input.RoleID = 2 // Assuming RoleID 2 is for admin

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = hashedPassword

	// Save admin to database
	if err := utils.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message or admin data
	c.JSON(http.StatusOK, gin.H{"message": "Admin added successfully", "admin": input})
}



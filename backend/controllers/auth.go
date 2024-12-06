package controllers

import (
	"net/http"
	"time"

	"pharmacy/database"
	"pharmacy/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("jwtkey")

type Credentials struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	ManufacturerName string `json:"manufacturerName"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing error"})
		return
	}

	// Check if a user with the same name exists
	var existingUser models.User
	if err := database.AuthDB.Where("username = ?", creds.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this username already exists"})
		return
	}

	// Check if a user with the same email exists
	if err := database.AuthDB.Where("email = ?", creds.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	// Get the "user" role
	var role models.Role
	if err := database.AuthDB.Where("rolename = ?", creds.Role).First(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user role"})
		return
	}

	addData := models.DataAdditional{
		ManufacturerName: creds.ManufacturerName,
	}

	// Create the DataAdditional record first
	if err := database.AuthDB.Create(&addData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating additional data"})
		return
	}

	user := models.User{
		Username:         creds.Username,
		PasswordHash:     string(hashedPassword),
		Email:            creds.Email,
		RoleID:           role.RoleID,
		DataAdditionalID: addData.DataAdditionalID,
	}
	if err := database.AuthDB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// Find the user by username
	var user models.User
	if err := database.AuthDB.Preload("DataAdditional").Preload("Role").Preload("Role.Permissions").Where("username = ?", creds.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	expirationTime := time.Now().Add(time.Hour * 72) // Token is valid for 72 hours
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // `jwtKey` should be secret key
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	// Send the token and permissions in the response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := database.AuthDB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Roles not found"})
		return
	}

	c.JSON(http.StatusAccepted, roles)
}

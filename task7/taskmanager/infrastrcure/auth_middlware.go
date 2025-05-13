package infrastrcure

import (
	"os"
	"strings"


	"github.com/joho/godotenv"

	"log"

	
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}
  jwtSecret := os.Getenv("JWT_SECRET")
  // here   we  are  going to have  the function as the validator
  
	return func(c*gin.Context){

		authheader:= c.GetHeader("Authorization")
		if authheader==" "{
			c.JSON(401,"authorization required")
}

authPart :=strings.Split(authheader," ")

if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
	c.JSON(401, gin.H{"error": "Invalid authorization header"})
	c.Abort()
	return


	}


	clamis, err :=  ValidateJWT(authPart[1] ,jwtSecret )
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	role, ok := clamis["role"].(string)
if !ok {
    c.JSON(403, gin.H{"error": "Role claim is missing or invalid"})
    c.Abort()
    return
}
c.Set("role", role)

	  
	}
}


func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.JSON(403, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		role, ok := roleValue.(string)
		if !ok || role != requiredRole {
			c.JSON(403, gin.H{"error": "Forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

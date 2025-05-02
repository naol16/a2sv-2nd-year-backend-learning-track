package middleware

import (
	"os"
	"strings"
	"github.com/joho/godotenv"
     "fmt"
	 "log"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc{
	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}
  jwtSecret := os.Getenv("JWT_SECRET")
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
	token, err := jwt.Parse(authPart[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	  
		return []byte(jwtSecret), nil
	  })
	  
	  if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return
	  }

	  if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if role, ok := claims["role"].(string); ok {
			c.Set("role", role)
		} else {
			c.JSON(403, gin.H{"error": "Role claim is missing or invalid"})
			c.Abort()
			return
		}
	} else {
		c.JSON(401, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
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
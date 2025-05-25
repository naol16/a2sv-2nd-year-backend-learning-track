package infrastrcure

import (
	
	"fmt"
	"os"
	"taskmanager/domain"
	

	"github.com/dgrijalva/jwt-go"
	
)



func ValidateJWT(tokenString string, secretKey string) (jwt.MapClaims, error) {

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid token: %v", err)
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims, nil
    }

    return nil, fmt.Errorf("unable to parse claims")
}



func   Generator(user  domain.User)  (jwttoken  string ,err error) {

	mysecretkey := os.Getenv("JWT_SECRET")
	fmt.Println(mysecretkey)
	var jwtsecret= []byte (mysecretkey)
	fmt.Println(user.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"name": user.Name,
		"role": string(user.Role),
	  })
	  
	  jwtToken, err := token.SignedString(jwtsecret)
	  return jwtToken,err



}
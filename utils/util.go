package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func Cleanup() {
//	var ctx = context.Background()
//
//	_, err := db.Exec(ctx, "DELETE FROM users")
//	if err != nil {
//		log.Fatal(err)
//	}
//}

package services
import "github.com/golang-jwt/jwt/v5"
import "time"

func GenerateToken(userID int, role string, secret string)(string,error)  {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	return token.SignedString([]byte(secret))
}
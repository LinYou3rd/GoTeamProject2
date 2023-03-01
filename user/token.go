package user

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("fjaklfjajfalkgh")

type Claims struct {
	ID        uint
	UserName  string
	Authority int `json:"authority"`
	jwt.StandardClaims
}

func GenerateTokenForUser(user userModel) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        user.ID,
		UserName:  user.Name,
		Authority: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "JUST",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, err
}

func GenerateTokenForAdmin(admin adminModel) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        admin.ID,
		Authority: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "JUST",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}

	return nil, err

}

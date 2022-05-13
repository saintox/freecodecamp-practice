package configs

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type (
	AuthService interface {
		GenerateTokenUser(token int) (string, error)
		ValidateToken(token string) (*jwt.Token, error)
	}

	jwtService struct {
		//
	}

	MyClaim struct {
		jwt.StandardClaims
		Email string `json:"email"`
	}
)

var (
	JwtSigningMethod = jwt.SigningMethodHS256
	JwtSignatureKey  = []byte("iloveindonesia")
)

func NewServiceAuth() *jwtService {
	return &jwtService{}
}

func (j jwtService) GenerateTokenUser(userId int) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = userId

	token := jwt.NewWithClaims(JwtSigningMethod, claim)

	signedToken, errJwt := token.SignedString(JwtSignatureKey)

	if errJwt != nil {
		return signedToken, errJwt
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, errJwt := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return JwtSignatureKey, nil
	})

	if errJwt != nil {
		return token, errJwt
	}

	return token, nil
}

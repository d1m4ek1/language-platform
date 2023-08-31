package sharedjwt

import (
	"english/backend/internal/services/svc-auth/models"
	"english/backend/shared/errorlog"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Login string
}

func (w *JwtWrapper) GenerateToken(client models.Client) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    client.Id,
		Login: client.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", errorlog.NewError(err.Error())
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (*jwtClaims, error) {
	claims := &jwtClaims{}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return claims, errorlog.NewError(err.Error())
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errorlog.NewError(err.Error())
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errorlog.NewError(err.Error())
	}

	return claims, nil

}

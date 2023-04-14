package helper

import (
	"Service-API/model"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func EncodeToken(adminId int64) (accessToken string, err error) {
	accessUUID := uuid.NewString()

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return accessToken, err
	}

	now := time.Now().UTC()

	exp := now.AddDate(0, 0, 30)

	claims := make(jwt.MapClaims)
	claims["idAdmin"] = adminId
	claims["exp"] = exp.Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix() // The time at which the token was issued.
	claims["iss"] = os.Getenv("JWT_ISS")
	claims["aud"] = os.Getenv("JWT_AUD")
	claims["accessUUID"] = accessUUID

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err = token.SignedString(key)
	if err != nil {
		return accessToken, err
	}

	redisKey := fmt.Sprintf("cmsv1-token-%s", accessUUID)

	InsertRedis(model.SetDataRedis{
		Key:  redisKey,
		Data: accessToken,
		Exp:  time.Hour * 730,
	})

	return accessToken, nil
}

func ValidateToken(encodedToken string) (token *jwt.Token, errData error) {
	jwtPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return token, err
	}

	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}

type DecodedToken struct {
	AdminId    string `json:"idAdmin"`
	AccessUUID string `json:"accessUUID"`
}

func DecodeToken(tokenString string) (decodedResult DecodedToken, err error) {
	jwtPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))

	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj DecodedToken
	if err := json.Unmarshal(jsonbody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}

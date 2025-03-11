package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/quietdevil/ChatSevice/cli-utils/models"
	"google.golang.org/grpc/metadata"
	"os"
)

func MarshalTokensInFile(anyStruct any, fileName string) error {
	jsonBytes, err := json.Marshal(anyStruct)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(jsonBytes)
	return err
}

func UnmarshalTokensInFile() (*models.Login, error) {
	login := &models.Login{}
	var buffer1 bytes.Buffer

	file, err := os.Open("login.json")
	if err != nil {
		return nil, err
	}

	_, err = buffer1.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(buffer1.String()), login)
	if err != nil {
		return nil, err
	}

	return login, nil
}

func UsernameFromAccessToken(token, secretKey string) (string, error) {
	claims, err := VerifyToken(token, []byte(secretKey))
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}

func VerifyToken(tokenStr string, secretKey []byte) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&models.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}
	return claims, nil
}

func NewContextOutGoing(ctx context.Context) (context.Context, error) {
	login, err := UnmarshalTokensInFile()
	if err != nil {
		return nil, err
	}
	md := metadata.New(map[string]string{"authorization": "Bearer " + login.AccessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, nil
}

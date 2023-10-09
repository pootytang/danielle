package helpers

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CREATE
func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	slog.Info("CreateToken(): Creating a token")

	slog.Info("CreateToken(): Decoding Private key")
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("CreateToken(): error decoding key %s", err.Error()))
		return "", fmt.Errorf("unable to decode key")
	}
	slog.Debug("CreateToken(): Decoded Private key successfully. Parsing Private key")

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		slog.Error(fmt.Sprintf("CreateToken(): parsing private key failed with: %s", err.Error()))
		return "", fmt.Errorf("CreateToken(): failed to parse key")
	}
	slog.Debug("CreateToken(): Parsed private key successfully")

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		slog.Error(fmt.Sprintf("CreateToken(): unable to sign token: %s", err.Error()))
		return "", fmt.Errorf("CreateToken(): failed to sign token")
	}

	slog.Info("CreateToken(): Access Token created")
	slog.Debug(fmt.Sprintf("CreateToken(): access token %s", token))
	return token, nil
}

// VALIDATE
func ValidateToken(token string, publicKey string) (interface{}, error) {
	slog.Info("ValidateToken(): Validating Token")

	slog.Info("ValidateToken(): Decoding Public key")
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		slog.Error(fmt.Sprintf("ValidateToken(): Error decoding public key: %s", err.Error()))
		return nil, fmt.Errorf("ValidateToken(): could not decode pub key")
	}
	slog.Debug("ValidateToken(): Decoded Public Key successfully. Parsing Private Key")

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		slog.Error(fmt.Sprintf("ValidateToken(): Error parsing decoded pub key: %s", err.Error()))
		return "", fmt.Errorf("validateToken(): parse pub key error")
	}
	slog.Debug("ValidateToken(): Parsed Private Key successfully")

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("ValidateToken->Parse(): unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		slog.Error(fmt.Sprintf("ValidateToken->Parse(): Error parsing token: %s", err.Error()))
		return nil, fmt.Errorf("validateToken->Parse(): parse error")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		slog.Error("ValidateToken(): problem with token, token is invalid")
		return nil, fmt.Errorf("ValidateToken(): invalid token")
	}

	slog.Info("ValidateToken(): Token Validation was successful")
	return claims["sub"], nil
}

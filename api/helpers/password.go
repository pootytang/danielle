package helpers

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	slog.Info("HashPassword(): Hashing password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		slog.Error(fmt.Sprintf("HashPassword(): error hashing the password: %s", err.Error()))
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, plainPassword string) error {
	slog.Info("VerifyPassword(): Called")
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

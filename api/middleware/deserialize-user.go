package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/helpers"
	"github.com/pootytang/danielleapi/initializers"
	"github.com/pootytang/danielleapi/models"
)

func DeserializeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		slog.Info("DeserializeUser->func(): Checking header and cookie for access token")

		// TRYING TO FIND THE ACCESS TOKEN
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		slog.Info("DeserializeUser->func(): Checking Authorization Header")
		authorizationHeader := ctx.Request().Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		slog.Debug(fmt.Sprintf("DeserializeUser->func(): Authorization Header set to %s", authorizationHeader))
		slog.Debug(fmt.Sprintf("DeserializeUser->func(): Authorization Header fields %s", fields))

		if len(fields) != 0 && fields[0] == "Bearer" {
			slog.Debug("DeserializeUser->func(): Bearer found in Authorization header, setting access token to the second field")
			access_token = fields[1]
		} else if err == nil {
			slog.Debug("DeserializeUser->func(): Grabbing access token from cookie")
			access_token = cookie.Value
		}

		if access_token == "" {
			slog.Error("DeserializeUser->func(): access_token not found")
			message := map[string]string{"status": "fail", "message": "not authorized"}
			// return echo.NewHTTPError(http.StatusUnauthorized, message)
			return ctx.JSON(http.StatusUnauthorized, message)
		}
		slog.Info("DeserializeUser->func(): access_token FOUND")
		slog.LogAttrs(
			context.Background(),
			slog.LevelDebug,
			"DeserializeUser->func(): Access Token retrieved",
			slog.String("value", access_token),
		)

		// VALIDATE THE TOKEN
		config, _ := initializers.LoadConfig(".")
		slog.LogAttrs(
			context.Background(),
			slog.LevelDebug,
			"DeserializeUser->func(): config loaded",
			slog.String("access_token_pubkey", config.AccessTokenPublicKey),
		)

		slog.Debug("DeserializeUser->func(): Validating access token")
		sub, err := helpers.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			slog.Error(fmt.Sprintf("DeserializeUser->func(): Error validating token: %s", err.Error()))
			message := map[string]string{"status": "fail", "message": "invalid access token"}
			// return echo.NewHTTPError(http.StatusUnauthorized, message)
			return ctx.JSON(http.StatusUnauthorized, message)
		}
		slog.Info("DeserializeUser->func(): access token is valid")

		// SEARCH FOR USER
		slog.Debug("DeserializeUser->func(): Searching for user in db")
		var user models.User
		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			slog.Error(fmt.Sprintf("DeserializeUser->func(): Error searching subject %s: %s", sub, err.Error()))
			message := map[string]string{"status": "fail", "message": "subject of token is not found in DB"}
			// return echo.NewHTTPError(http.StatusUnauthorized, message)
			return ctx.JSON(http.StatusForbidden, message)
		}

		slog.Info("DeserializeUser->func(): adding user to context")
		ctx.Set("currentUser", user)
		return next(ctx)
	}
}

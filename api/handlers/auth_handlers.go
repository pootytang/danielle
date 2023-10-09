package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/helpers"
	"github.com/pootytang/danielleapi/initializers"
	"github.com/pootytang/danielleapi/models"
	"github.com/pootytang/danielleapi/types"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

/********** REGISTER USER **********/
func (ac *AuthController) RegisterUser(ctx echo.Context) error {
	slog.Info("RegisterUser(): Signing up new user")
	var payload *models.SignUpInput

	slog.Debug("RegisterUser(): Binding to post data")
	if err := ctx.Bind(&payload); err != nil {
		slog.Error(fmt.Sprintf("RegisterUser(): Error binding to payload: %s", err.Error()))
		message := map[string]string{"status": "error", "message": "signup user bind error"}
		return ctx.JSON(http.StatusBadRequest, message)
	}

	slog.Debug("RegisterUser(): Confirming password")
	if payload.Password != payload.PasswordConfirm {
		slog.Error("RegisterUser(): Passwords do not match")
		message := map[string]string{"status": "error", "message": "password mismatch"}
		return ctx.JSON(http.StatusBadRequest, message)
	}

	slog.Debug("RegisterUser(): Hashing password")
	hashedPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		slog.Error(fmt.Sprintf("RegisterUser(): Error hashing the password: %s", err.Error()))
		message := map[string]string{"status": "error", "message": "password hashing error"}
		return ctx.JSON(http.StatusBadGateway, message)
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}

	slog.Info(("RegisterUser(): Checking if user is delane or angel"))
	if strings.ToLower(payload.Email) == "delane.jackson@gmail.com" || strings.ToLower(payload.Email) == "angelica.t.a.jackson@gmail.com" {
		slog.Warn("RegisterUser(): new user is Delane or Angel. Setting Role to admin")
		newUser.Role = "admin"
	}

	slog.LogAttrs(context.Background(), slog.LevelInfo, "RegisterUser(): Creating new user",
		slog.String("name", newUser.Name),
		slog.String("email", newUser.Email),
		slog.String("role", newUser.Role),
	)
	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		slog.Warn(fmt.Sprintf("RegisterUser(): User may exist: %s", result.Error.Error()))
		message := map[string]string{"status": "warn", "message": "user exists"}
		return ctx.JSON(http.StatusConflict, message)
	} else if result.Error != nil {
		slog.Error(fmt.Sprintf("RegisterUser(): Error creating user: %s", result.Error.Error()))
		message := map[string]string{"status": "error", "message": "unable to create user"}
		return ctx.JSON(http.StatusBadGateway, message)
	}
	slog.Info("RegisterUser(): User created successfully")

	slog.Debug("RegisterUser(): Loading config")
	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	slog.Info("RegisterUser(): Generating Access Token")
	access_token, err := helpers.CreateToken(config.AccessTokenExpiresIn, newUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("RegisterUser(): Error generating access token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "access token generation error"}
		return ctx.JSON(http.StatusBadGateway, message)
	}

	slog.Info("RegisterUser(): Generating Refresh Token")
	refresh_token, err := helpers.CreateToken(config.RefreshTokenExpiresIn, newUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("RegisterUser(): Error generating refresh token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "refresh token generation error"}
		return ctx.JSON(http.StatusBadGateway, message)
	}

	// COOKIES
	// Removed this as I don't think I need it. Just send the tokens and logged in status in the json payload
	// Protected sites will look for this token in either a cookie returned from the app or authorization bearer header

	// _, cookieDomain := types.GetHost()
	// slog.Info("RegisterUser(): creating cookies")
	// access_token_cookie := new(http.Cookie)
	// access_token_cookie.Name = "access_token"
	// access_token_cookie.Value = access_token
	// access_token_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// access_token_cookie.Path = "/"
	// access_token_cookie.Domain = cookieDomain
	// access_token_cookie.Secure = types.COOKIE_SECURE
	// access_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(access_token_cookie)
	// slog.Debug(fmt.Sprintf("RegisterUser(): access_token cookie: %s", access_token_cookie))

	// refresh_token_cookie := new(http.Cookie)
	// refresh_token_cookie.Name = "refresh_token"
	// refresh_token_cookie.Value = refresh_token
	// refresh_token_cookie.MaxAge = config.RefreshTokenMaxAge * 60 * 24 * 5 // 5 days
	// refresh_token_cookie.Path = "/"
	// refresh_token_cookie.Domain = cookieDomain
	// refresh_token_cookie.Secure = types.COOKIE_SECURE
	// refresh_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(refresh_token_cookie)
	// slog.Debug(fmt.Sprintf("RegisterUser(): refresh_token cookie: %s", refresh_token_cookie))

	// logged_in_cookie := new(http.Cookie)
	// logged_in_cookie.Name = "logged_in"
	// logged_in_cookie.Value = "true"
	// logged_in_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// logged_in_cookie.Path = "/"
	// logged_in_cookie.Domain = cookieDomain
	// logged_in_cookie.Secure = types.COOKIE_SECURE
	// logged_in_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(logged_in_cookie)
	// slog.Debug(fmt.Sprintf("RegisterUser(): logged_in cookie: %s", logged_in_cookie))

	// slog.Info("RegisterUser(): Cookies have been set")

	slog.Info(fmt.Sprintf("RegisterUser(): User %s is created and logged in. Creating json response", newUser.Name))
	userResponse := &models.UserResponse{
		ID:            newUser.ID,
		Name:          newUser.Name,
		Email:         newUser.Email,
		Role:          newUser.Role,
		Access_Token:  access_token,
		Refresh_Token: refresh_token,
		Logged_In:     true,
		CreatedAt:     newUser.CreatedAt,
		UpdatedAt:     newUser.UpdatedAt,
	}

	message := map[string]interface{}{"status": "susccess", "message": "user created and logged in", "user": userResponse}
	return ctx.JSON(http.StatusCreated, message)
}

/********** SIGNIN USER **********/
func (ac *AuthController) SignInUser(ctx echo.Context) error {
	slog.Info("SignInUser(): Signing in user")
	var payload *models.SignInInput

	slog.Info("SignInUser(): Binding to post data")
	if err := ctx.Bind(&payload); err != nil {
		slog.Error(fmt.Sprintf("SignInUser(): Error binding to payload: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "user bind error"}
		return ctx.JSON(http.StatusBadRequest, message)
	}

	slog.Info("SignInUser(): checking DB for user")
	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		slog.Error(fmt.Sprintf("SignInUser(): Error searching DB for email %s: %s", payload.Email, result.Error.Error()))
		message := map[string]string{"status": "fail", "message": fmt.Sprintf("email %s not found", payload.Email)}
		return ctx.JSON(http.StatusUnauthorized, message)
	}

	slog.Info("SignInUser(): verifying users password")
	if err := helpers.VerifyPassword(user.Password, payload.Password); err != nil {
		slog.Error("SignInUser(): password verification failed")
		message := map[string]string{"status": "fail", "message": "password verification failed"}
		return ctx.JSON(http.StatusUnauthorized, message)
	}

	slog.Info("SignInUser(): Loading config")
	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	slog.Info("SignInUser(): Generating Access Token")
	access_token, err := helpers.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("SignInUser(): Error generating access token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "access token generation error"}
		return ctx.JSON(http.StatusInternalServerError, message)
	}

	slog.Debug("SignInUser(): Generating Refresh Token")
	refresh_token, err := helpers.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("SignInUser(): Error generating refresh token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "refresh token generation error"}
		return ctx.JSON(http.StatusInternalServerError, message)
	}

	// COOKIES
	// Removed this as I don't think I need it. Just send the tokens and logged in status in the json payload
	// Protected sites will look for this token in either a cookie returned from the app or authorization bearer header
	// _, cookieDomain := types.GetHost()
	// slog.Debug("SignInUser(): creating cookies")
	// access_token_cookie := new(http.Cookie)
	// access_token_cookie.Name = "access_token"
	// access_token_cookie.Value = access_token
	// access_token_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// access_token_cookie.Path = "/"
	// access_token_cookie.Domain = cookieDomain
	// access_token_cookie.Secure = types.COOKIE_SECURE
	// access_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(access_token_cookie)
	// slog.Debug(fmt.Sprintf("SignInUser(): access_token cookie: %s", access_token_cookie))

	// refresh_token_cookie := new(http.Cookie)
	// refresh_token_cookie.Name = "refresh_token"
	// refresh_token_cookie.Value = refresh_token
	// refresh_token_cookie.MaxAge = config.RefreshTokenMaxAge * 60
	// refresh_token_cookie.Path = "/"
	// refresh_token_cookie.Domain = cookieDomain
	// refresh_token_cookie.Secure = types.COOKIE_SECURE
	// refresh_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(refresh_token_cookie)
	// slog.Debug(fmt.Sprintf("SignInUser(): refresh_token cookie: %s", refresh_token_cookie))

	// logged_in_cookie := new(http.Cookie)
	// logged_in_cookie.Name = "logged_in"
	// logged_in_cookie.Value = "true"
	// logged_in_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// logged_in_cookie.Path = "/"
	// logged_in_cookie.Domain = cookieDomain
	// logged_in_cookie.Secure = types.COOKIE_SECURE
	// logged_in_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(logged_in_cookie)
	// slog.Debug(fmt.Sprintf("SignInUser(): logged_in cookie: %s", logged_in_cookie))

	// slog.Info("SignInUser(): Cookies have been set")

	slog.Info(fmt.Sprintf("SignInUser(): User %s has signed in", user.Name))
	userResponse := &models.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Role:          user.Role,
		Access_Token:  access_token,
		Refresh_Token: refresh_token,
		Logged_In:     true,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
	message := map[string]interface{}{"status": "success", "message": "logged in", "user": userResponse}
	return ctx.JSON(http.StatusOK, message)
}

/********** REFRESH TOKEN **********/
func (ac *AuthController) RefreshAccessToken(ctx echo.Context) error {
	slog.Info("RefreshToken(): Refreshing the access token")

	var rt *models.RefreshToken
	if err := ctx.Bind(&rt); err != nil {
		slog.Error(fmt.Sprintf("SignInUser(): Error binding to payload: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "user bind error"}
		return ctx.JSON(http.StatusBadRequest, message)
	}

	slog.Info("RefreshToken(): Retrieved refresh token")
	slog.Debug(fmt.Sprintf("RefreshToken(): refresh token: %s", rt.Refresh_Token))

	slog.Debug("RefreshToken(): loading config")
	config, _ := initializers.LoadConfig(".")

	sub, err := helpers.ValidateToken(rt.Refresh_Token, config.RefreshTokenPublicKey)
	if err != nil {
		slog.Error(fmt.Sprintf("RefreshToken():: Error validating refresh token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "refresh token validation error"}
		return ctx.JSON(http.StatusForbidden, message)
	}
	slog.Info("RefreshToken(): Refresh Token validation successful")

	slog.Debug(fmt.Sprintf("RefreshToken(): Checking DB for subject %s", sub))
	var user models.User
	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		slog.Error(fmt.Sprintf("RefreshToken():: User not found: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "user belonging to token was not found"}
		return ctx.JSON(http.StatusForbidden, message)
	}
	slog.Info(fmt.Sprintf("RefreshToken(): found user with email %s", user.Email))

	slog.Info("RefreshToken(): creating new access token")
	access_token, err := helpers.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("RefreshToken():: Access Token generation error: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "Unable to generate access token"}
		return ctx.JSON(http.StatusForbidden, message)
	}
	slog.Info(fmt.Sprintf("RefreshToken(): Access Token created: %s", access_token))

	// COOKIES
	// Removed this as I don't think I need it. Just send the tokens and logged in status in the json payload
	// Protected sites will look for this token in either a cookie returned from the app or authorization bearer header
	// _, cookieDomain := types.GetHost()
	// slog.Debug("RefreshToken(): creating cookies")
	// access_token_cookie := new(http.Cookie)
	// access_token_cookie.Name = "access_token"
	// access_token_cookie.Value = access_token
	// access_token_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// access_token_cookie.Path = "/"
	// access_token_cookie.Domain = cookieDomain
	// access_token_cookie.Secure = types.COOKIE_SECURE
	// access_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(access_token_cookie)
	// slog.Debug(fmt.Sprintf("RefreshToken(): access_token cookie: %s", access_token_cookie))

	// logged_in_cookie := new(http.Cookie)
	// logged_in_cookie.Name = "logged_in"
	// logged_in_cookie.Value = "true"
	// logged_in_cookie.MaxAge = config.AccessTokenMaxAge * 60
	// logged_in_cookie.Path = "/"
	// logged_in_cookie.Domain = cookieDomain
	// logged_in_cookie.Secure = types.COOKIE_SECURE
	// logged_in_cookie.HttpOnly = types.COOKIE_HTTPONLY
	// ctx.SetCookie(logged_in_cookie)
	// slog.Debug(fmt.Sprintf("RefreshToken(): logged_in cookie: %s", logged_in_cookie))

	// slog.Info("RefreshToken(): Cookies have been set")

	slog.Info("RefreshToken(): Access Token has been refreshed")
	message := map[string]string{"status": "success", "message": "new access token generated", "access_token": access_token}
	return ctx.JSON(http.StatusOK, message)
}

/********** GET USER **********/
func (ac *AuthController) GetUser(ctx echo.Context) error {
	slog.Info("GetUser(): Retrieving User")
	var payload *models.FindUserByAccessToken

	slog.Info("GetUser(): Binding to post data")
	if err := ctx.Bind(&payload); err != nil {
		slog.Error(fmt.Sprintf("GetUser(): Error binding to payload: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "token bind error"}
		return ctx.JSON(http.StatusBadRequest, message)
	}
	slog.Debug("GetUser(): Retrieved Access Token")

	// VALIDATE THE TOKEN
	config, _ := initializers.LoadConfig(".")
	slog.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		"GetUser(): config loaded",
		slog.String("access_token_pubkey", config.AccessTokenPublicKey),
		slog.String("access_token", payload.Access_Token),
	)

	slog.Info("GetUser(): Validating access token")
	sub, err := helpers.ValidateToken(payload.Access_Token, config.AccessTokenPublicKey)
	if err != nil {
		slog.Error(fmt.Sprintf("GetUser(): Error validating token: %s", err.Error()))
		message := map[string]string{"status": "fail", "message": "invalid access token"}
		return ctx.JSON(http.StatusUnauthorized, message)
	}
	slog.Info("GetUser(): access token is valid")

	// SEARCH FOR USER
	slog.Info("GetUser(): Searching for user in db")
	var user models.User
	result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		slog.Error(fmt.Sprintf("GetUser(): Error searching subject %s: %s", sub, err.Error()))
		message := map[string]string{"status": "fail", "message": "subject of token is not found in DB"}
		return ctx.JSON(http.StatusForbidden, message)
	}
	slog.Info(fmt.Sprintf("GetUser(): Found User %s ", user.Name))
	userResponse := &models.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Role:          user.Role,
		Access_Token:  payload.Access_Token,
		Refresh_Token: "",
		Logged_In:     true,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
	message := map[string]interface{}{"status": "success", "message": "logged in", "user": userResponse}
	return ctx.JSON(http.StatusOK, message)
}

/********** LOGOUT USER **********/
// Not using cookies anymore so this endpoint may not be needed. Can logout clientside
func (ac *AuthController) LogoutUser(ctx echo.Context) error {
	slog.Info("LogoutUser(): Logging user out")

	slog.Debug("LogoutUser(): nullifying access_token cookie")
	_, cookieDomain := types.GetHost()
	access_token_cookie := new(http.Cookie)
	access_token_cookie.Name = "access_token"
	access_token_cookie.Value = ""
	access_token_cookie.MaxAge = -1
	access_token_cookie.Path = "/"
	access_token_cookie.Domain = cookieDomain
	access_token_cookie.Secure = types.COOKIE_SECURE
	access_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	ctx.SetCookie(access_token_cookie)

	slog.Debug("LogoutUser(): nullifying refresh_token cookie")
	refresh_token_cookie := new(http.Cookie)
	refresh_token_cookie.Name = "refresh_token"
	refresh_token_cookie.Value = ""
	refresh_token_cookie.MaxAge = -1
	refresh_token_cookie.Path = "/"
	refresh_token_cookie.Domain = cookieDomain
	refresh_token_cookie.Secure = types.COOKIE_SECURE
	refresh_token_cookie.HttpOnly = types.COOKIE_HTTPONLY
	ctx.SetCookie(refresh_token_cookie)

	slog.Debug("LogoutUser(): nullifying logged_in cookie")
	logged_in_cookie := new(http.Cookie)
	logged_in_cookie.Name = "logged_in"
	logged_in_cookie.Value = ""
	logged_in_cookie.MaxAge = -1
	logged_in_cookie.Path = "/"
	logged_in_cookie.Domain = cookieDomain
	logged_in_cookie.Secure = types.COOKIE_SECURE
	logged_in_cookie.HttpOnly = types.COOKIE_HTTPONLY
	ctx.SetCookie(logged_in_cookie)

	slog.Info("LogoutUser(): logged out successfully")
	message := map[string]string{"status": "success", "message": "logged out"}
	return ctx.JSON(http.StatusOK, message)
}

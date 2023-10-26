package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/helpers"
	"github.com/pootytang/danielleapi/types"
)

func ShowAllRoutes(c echo.Context) error {
	route := types.Route{}
	routes := types.AllRoutes{
		PublicRoutes:  route.GetPublicRoutes(),
		PrivateRoutes: route.GetPrivateRoutes(),
	}

	return c.JSON(http.StatusOK, routes)
}

func ShowPublicRoutes(c echo.Context) error {
	route := types.Route{}
	return c.JSON(http.StatusOK, route.GetPublicRoutes())
}

/**** ULTRASOUND ****/
func GetUltrasounds(c echo.Context) error {
	slog.Info("GetUltrasounds(): Called")

	dir := types.ULTRASOUNDS_FOLDER
	route := types.ULTRASOUNDS_ROUTE

	files, err := helpers.GetFiles(dir, route)

	slog.Debug(fmt.Sprintf("GetUltraSounds(): Files == %s", files))
	if err != nil {
		slog.Error(fmt.Sprintf("GetUltraSounds(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	return c.JSON(http.StatusOK, files)
}

func GetUltrasoundImage(c echo.Context) error {
	slog.Info("GetUltrasoundImage(): Called")

	dir := types.ULTRASOUNDS_FOLDER
	file_name := c.Param("image")

	slog.Info("GetUltrasoundImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetUltrasoundImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetUltrasoundImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** GROWING MOMMY ****/
func GetGrowingMommyImages(c echo.Context) error {
	slog.Info("GetGrowingMommyImages(): Called")

	dir := types.GROWING_FOLDER
	route := types.GROWING_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	return c.JSON(http.StatusOK, files)
}

func GetGrowingMommyImage(c echo.Context) error {
	slog.Info("GetGrowingMommyImage(): Called")

	dir := types.GROWING_FOLDER
	file_name := c.Param("image")

	slog.Info("GetGrowingMommyImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetGrowingMommyImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetGrowingMommyImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** BABY SHOWER / REVEAL ****/
func GetShowerRevealImages(c echo.Context) error {
	slog.Info("GetShowerRevealImages(): Called")

	dir := types.SHOWERREVEAL_FOLDER
	route := types.SHOWERREVEAL_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	return c.JSON(http.StatusOK, files)
}

func GetShowerRevealImage(c echo.Context) error {
	slog.Info("GetShowerRevealImage(): Called")

	dir := types.SHOWERREVEAL_FOLDER
	file_name := c.Param("image")

	slog.Info("GetShowerRevealImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetShowerRevealImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetShowerRevealImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** BIRTH DAY ****/
func GetBirthDayImages(c echo.Context) error {
	slog.Info("GetBirthDayImages(): Called")

	dir := types.BIRTHDAY_FOLDER
	route := types.BIRTHDAY_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetBirthDayImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	return c.JSON(http.StatusOK, files)
}

func GetBirthDayImage(c echo.Context) error {
	slog.Info("GetBirthDayImage(): Called")

	dir := types.BIRTHDAY_FOLDER
	file_name := c.Param("image")

	slog.Info("GetBirthDayImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetBirthDayImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetBirthDayImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/helpers"
	"github.com/pootytang/danielleapi/types"
)

/**** INFANCY STAGE ****/
func GetInfantImages(c echo.Context) error {
	slog.Info("GetInfantImages(): Called")

	dir := types.INFANT_STAGE_FOLDER
	route := types.INFANT_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetInfantImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetInfantImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetInfantImage(c echo.Context) error {
	slog.Info("GetInfantImage(): Called")

	dir := types.INFANT_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetInfantImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetInfantImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetInfantImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** TODDLER STAGE ****/
func GetToddlerImages(c echo.Context) error {
	slog.Info("GetToddlerImages(): Called")

	dir := types.TODDLER_STAGE_FOLDER
	route := types.TODDLER_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetToddlerImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetToddlerImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetToddlerImage(c echo.Context) error {
	slog.Info("GetToddlerImage(): Called")

	dir := types.TODDLER_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetToddlerImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetToddlerImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetToddlerImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** EARLY CHILDHOOD STAGE ****/
func GetEarlyChildhoodImages(c echo.Context) error {
	slog.Info("GetEarlyChildhoodImages(): Called")

	dir := types.EARLYCHILDHOOD_STAGE_FOLDER
	route := types.EARLYCHILDHOOD_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetEarlyChildhoodImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetEarlyChildhoodImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetEarlyChildhoodImage(c echo.Context) error {
	slog.Info("GetEarlyChildhoodImage(): Called")

	dir := types.EARLYCHILDHOOD_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetEarlyChildhoodImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetEarlyChildhoodImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetEarlyChildhoodImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** LATE CHILDHOOD STAGE ****/
func GetLateChildhoodImages(c echo.Context) error {
	slog.Info("GetLateChildhoodImages(): Called")

	dir := types.LATECHILDHOOD_STAGE_FOLDER
	route := types.LATECHILDHOOD_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetLateChildhoodImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetLateChildhoodImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetLateChildhoodImage(c echo.Context) error {
	slog.Info("GetLateChildhoodImage(): Called")

	dir := types.LATECHILDHOOD_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetLateChildhoodImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetLateChildhoodImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetLateChildhoodImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** ADOLESCENCE STAGE ****/
func GetAdolescenceImages(c echo.Context) error {
	slog.Info("GetAdolescenceImages(): Called")

	dir := types.ADOLESCENCE_STAGE_FOLDER
	route := types.ADOLESCENCE_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetAdolescenceImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetAdolescenceImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetAdolescenceImage(c echo.Context) error {
	slog.Info("GetAdolescenceImage(): Called")

	dir := types.ADOLESCENCE_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetAdolescenceImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetAdolescenceImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetAdolescenceImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** EARLY ADULTHOOD STAGE ****/
func GetEarlyAdulthoodImages(c echo.Context) error {
	slog.Info("GetEarlyAdulthoodImages(): Called")

	dir := types.EARLYADULTHOOD_STAGE_FOLDER
	route := types.EARLYADULTHOOD_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetEarlyAdulthoodImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetEarlyAdulthoodImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetEarlyAdulthoodImage(c echo.Context) error {
	slog.Info("GetEarlyAdulthoodImage(): Called")

	dir := types.EARLYADULTHOOD_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetEarlyAdulthoodImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetEarlyAdulthoodImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetEarlyAdulthoodImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** MIDDLE ADULTHOOD STAGE ****/
func GetMiddleAdulthoodImages(c echo.Context) error {
	slog.Info("GetMiddleAdulthoodImages(): Called")

	dir := types.MIDDLEADULTHOOD_STAGE_FOLDER
	route := types.MIDDLEADULTHOOD_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetMiddleAdulthoodImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetMiddleAdulthoodImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetMiddleAdulthoodImage(c echo.Context) error {
	slog.Info("GetMiddleAdulthoodImage(): Called")

	dir := types.MIDDLEADULTHOOD_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetMiddleAdulthoodImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetMiddleAdulthoodImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetMiddleAdulthoodImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

/**** LATE ADULTHOOD STAGE ****/
func GetLateAdulthoodImages(c echo.Context) error {
	slog.Info("GetLateAdulthoodImages(): Called")

	dir := types.LATEADULTHOOD_STAGE_FOLDER
	route := types.LATEADULTHOOD_STAGE_ROUTE

	files, err := helpers.GetFiles(dir, route)

	if err != nil {
		slog.Error(fmt.Sprintf("GetLateAdulthoodImages(): Error retrieving files: %s", error.Error(err)))
		message := map[string]string{"message": "Folder not found"}
		return c.JSON(http.StatusNotFound, message)
	}

	slog.Info(fmt.Sprintf("GetLateAdulthoodImages(): Returning %d images", len(files)))
	return c.JSON(http.StatusOK, files)
}

func GetLateAdulthoodImage(c echo.Context) error {
	slog.Info("GetLateAdulthoodImage(): Called")

	dir := types.LATEADULTHOOD_STAGE_FOLDER
	file_name := c.Param("image")

	slog.Info("GetLateAdulthoodImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetLateAdulthoodImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetLateAdulthoodImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

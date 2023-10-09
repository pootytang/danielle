package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/helpers"
	"github.com/pootytang/danielleapi/models"
	"github.com/pootytang/danielleapi/types"
)

/***** ALL ROUTES *****/
func ShowPrivateRoutes(c echo.Context) error {
	route := types.Route{}
	return c.JSON(http.StatusOK, route.GetPrivateRoutes())
}

/***** DELIVERY *****/
func GetDeliveryImages(c echo.Context) error {
	slog.Info("GetDeliveryImages(): Called")

	slog.Info("GetDeliveryImages(): Checking the context for user")
	current_user := c.Get("currentUser").(models.User)
	slog.Debug(fmt.Sprintf("GetDeliveryImages(): pulled %s from context with role %s", current_user.Name, current_user.Role))

	if (current_user == models.User{}) || (current_user.Role != "admin") {
		slog.Warn("GetDeliveryImages(): no user found or invalid role")
		message := map[string]string{"message": "admin user required"}
		return c.JSON(http.StatusForbidden, message)
	}

	slog.Info(fmt.Sprintf("GetDeliveryImages(): User %s is an admin user. Getting images", current_user.Name))

	dir := types.PRIVATE_DELIVERY_FOLDER
	route := types.DELIVERY_ROUTE

	files, err := helpers.GetFiles(dir, route)

	slog.Debug(fmt.Sprintf("GetDeliveryImages(): Files == %s", files))
	if err != nil {
		slog.Error(fmt.Sprintf("GetDeliveryImages(): Error retrieving images: %s", err.Error()))
		message := map[string]string{"message": "error retrieving files"}
		return c.JSON(http.StatusServiceUnavailable, message)
	}

	slog.Info("GetDeliveryImages(): Retrieved images successfully")
	return c.JSON(http.StatusOK, files)
}

func GetDeliveryImage(c echo.Context) error {
	slog.Info("GetDeliveryImage(): Called")

	/**** CHECK IF USER IS AUTHENTICATED ****/
	slog.Info("GetDeliveryImage(): Checking the context for user")
	current_user := c.Get("currentUser").(models.User)
	slog.Debug(fmt.Sprintf("GetDeliveryImage(): pulled %s from context with role %s", current_user.Name, current_user.Role))

	if (current_user == models.User{}) || (current_user.Role != "admin") {
		slog.Warn("GetDeliveryImage(): no user found or invalid role")
		message := map[string]string{"message": "admin user required"}
		return c.JSON(http.StatusForbidden, message)
	}

	slog.Info(fmt.Sprintf("GetDeliveryImage(): User %s is an admin user. Getting image", current_user.Name))

	/**** GET THE REQUESTED IMAGE ****/
	dir := types.PRIVATE_DELIVERY_FOLDER
	file_name := c.Param("image")

	slog.Info("GetDeliveryImage(): Checking image " + dir + file_name)
	if helpers.FileExists(dir + file_name) {
		slog.Debug("GetDeliveryImage(): FOUND " + file_name)
		return c.File(dir + file_name)
	} else {
		slog.Warn(fmt.Sprintf("GetDeliveryImage(): Image %s doesn't exist", dir+file_name))
		message := map[string]string{"message": file_name + " not found"}
		return c.JSON(http.StatusNotFound, message)
	}
}

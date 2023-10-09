package types

import (
	"os"
	"strings"
)

type Route struct {
	Name         string `json:"name"`
	FilePath     string `json:"file_path,omitempty"`
	URI          string `json:"uri"`
	Description  string `json:"description"`
	Authrequired bool   `json:"auth_required"`
}

type AllRoutes struct {
	PublicRoutes  []Route
	PrivateRoutes []Route
}

// Used in fshelp.go but placed here for single config location
func GetHost() (url string, hostname string) {
	url = os.Getenv("PUBLIC_HOST_URL")
	urlArr := strings.Split(url, "://")
	host := strings.Split(urlArr[1], ":")[0]
	// var host string

	// switch os.Getenv("ENV") {
	// case "PROD": // Running in a container on Synology
	// 	host = "https://dani-belle.com:4443"
	// case "DEV": // Running in a container locally (need to edit host file to point the name to the loopback)
	// 	host = "https://dani-belle.com:4443"
	// case "LOCAL": // No container involved
	// 	host = "http://localhost:1323"
	// default:
	// 	slog.Warn("GetHost(): Environment Var not found, Setting to default host value")
	// 	host = "http://localhost:1323"
	// }

	// return host

	return url, host
}

func (r *Route) GetPublicRoutes() []Route {
	publicroutes := []Route{
		{
			Name:         "index",
			URI:          "/",
			Description:  "displays a list of routes",
			Authrequired: false,
		},
		{
			Name:         "Ultrasounds",
			URI:          ULTRASOUNDS_ROUTE,
			Description:  "pics of baby while in mommy's tummy",
			Authrequired: false,
		},
		{
			Name:         "Growing Mommy",
			URI:          GROWING_ROUTE,
			Description:  "pics of mommy showing her tummy",
			Authrequired: false,
		},
		{
			Name:         "Day of Birth",
			URI:          BIRTHDAY_ROUTE,
			Description:  "Day of Birth pics",
			Authrequired: false,
		},
	}

	return publicroutes
}

func (r *Route) GetPrivateRoutes() []Route {
	privateroutes := []Route{
		{
			Name:         "Delivery",
			URI:          DELIVERY_ROUTE,
			Description:  "Images and Vids of the actual delivery",
			Authrequired: true,
		},
	}

	return privateroutes
}

package types

import (
	"os"
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

const (
	LOG_FILE string = "/var/log/api/api.log"

	/**** BASE FOLDER CONSTANTS ****/
	PUBLIC_FOLDER  = "/media/public"
	PRIVATE_FOLDER = "/media/private"

	PREBIRTH_ROOT_FOLDER = PUBLIC_FOLDER + "/prebirth"
	BIRTH_ROOT_FOLDER    = PUBLIC_FOLDER + "/birth"

	/**** PREBIRTH FOLDERs CONSTANTS ****/
	ULTRASOUNDS_FOLDER  = PREBIRTH_ROOT_FOLDER + "/ultrasounds/"
	GROWING_FOLDER      = PREBIRTH_ROOT_FOLDER + "/growing_mommy/"
	SHOWERREVEAL_FOLDER = PREBIRTH_ROOT_FOLDER + "/shower_reveal/"

	/**** BIRTH FOLDERs CONSTANTS ****/
	BIRTHDAY_FOLDER = BIRTH_ROOT_FOLDER + "/birthday/"

	/**** PRIVATE DELIVERY FOLDERs CONSTANTS ****/
	PRIVATE_DELIVERY_FOLDER = PRIVATE_FOLDER + "/birth/delivery/"

	/**** PRIVATE BIRTH FOLDERs CONSTANS ****/

	/**** ROUTE CONSTANTS ****/
	public_route       = "/api/v1/public"
	private_route      = "/api/v1/private"
	ULTRASOUNDS_ROUTE  = public_route + "/ultrasounds/"
	GROWING_ROUTE      = public_route + "/growingmommy/"
	SHOWERREVEAL_ROUTE = public_route + "/shower_reveal/"
	BIRTHDAY_ROUTE     = public_route + "/dob/"
	DELIVERY_ROUTE     = private_route + "/delivery/"
)

// Used to map the endpoint in a request to the matching folder on the filesystem. Some are the same but some are not
var ENDPOINT_TO_FOLDER = map[string]string{
	"ultrasounds":   "ultrasounds",
	"growingmommy":  "growing_mommy",
	"shower_reveal": "shower_reveal",
}

// Used in fshelp.go but placed here for single config location
func GetHost() string {
	return os.Getenv("PUBLIC_HOST_URL")
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

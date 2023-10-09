package types

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

	/**** COOKIE VALUES ****/
	COOKIE_SECURE   bool = false
	COOKIE_HTTPONLY      = false
)

// Used to map the endpoint in a request to the matching folder on the filesystem. Some are the same but some are not
var ENDPOINT_TO_FOLDER = map[string]string{
	"ultrasounds":   "ultrasounds",
	"growingmommy":  "growing_mommy",
	"shower_reveal": "shower_reveal",
	"delivery":      "delivery",
}

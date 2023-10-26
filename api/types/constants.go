package types

const (
	LOG_FILE string = "/var/log/api/api.log"

	/**** BASE FOLDER CONSTANTS ****/
	PUBLIC_FOLDER  = "/media/public"
	PRIVATE_FOLDER = "/media/private"

	PREBIRTH_ROOT_FOLDER = PUBLIC_FOLDER + "/prebirth"
	BIRTH_ROOT_FOLDER    = PUBLIC_FOLDER + "/birth"
	STAGES_ROOT_FOLDER   = PUBLIC_FOLDER + "/stages"

	/**** PREBIRTH FOLDERs CONSTANTS ****/
	ULTRASOUNDS_FOLDER  = PREBIRTH_ROOT_FOLDER + "/ultrasounds/"
	GROWING_FOLDER      = PREBIRTH_ROOT_FOLDER + "/growing_mommy/"
	SHOWERREVEAL_FOLDER = PREBIRTH_ROOT_FOLDER + "/shower_reveal/"

	/**** BIRTH FOLDERs CONSTANTS ****/
	BIRTHDAY_FOLDER = BIRTH_ROOT_FOLDER + "/day_of_birth/"

	/**** ROUTE CONSTANTS ****/
	public_route       = "/api/v1/public"
	private_route      = "/api/v1/private"
	ULTRASOUNDS_ROUTE  = public_route + "/ultrasounds/"
	GROWING_ROUTE      = public_route + "/growingmommy/"
	SHOWERREVEAL_ROUTE = public_route + "/shower_reveal/"
	BIRTHDAY_ROUTE     = public_route + "/dob/"

	/**** STAGES CONSTANTS ****/
	INFANT_STAGE_FOLDER          = STAGES_ROOT_FOLDER + "/zero_to_one/"
	INFANT_STAGE_ROUTE           = public_route + "/zero-one/"
	TODDLER_STAGE_FOLDER         = STAGES_ROOT_FOLDER + "/one_to_two/"
	TODDLER_STAGE_ROUTE          = public_route + "/one-two/"
	EARLYCHILDHOOD_STAGE_FOLDER  = STAGES_ROOT_FOLDER + "/three_to_six/"
	EARLYCHILDHOOD_STAGE_ROUTE   = public_route + "/three-six/"
	LATECHILDHOOD_STAGE_FOLDER   = STAGES_ROOT_FOLDER + "/seven_to_ten/"
	LATECHILDHOOD_STAGE_ROUTE    = public_route + "/seven-ten/"
	ADOLESCENCE_STAGE_FOLDER     = STAGES_ROOT_FOLDER + "/eleven_to_nineteen/"
	ADOLESCENCE_STAGE_ROUTE      = public_route + "/eleven-nineteen/"
	EARLYADULTHOOD_STAGE_FOLDER  = STAGES_ROOT_FOLDER + "/twenty_to_fortyfour/"
	EARLYADULTHOOD_STAGE_ROUTE   = public_route + "/twenty-fortyfour/"
	MIDDLEADULTHOOD_STAGE_FOLDER = STAGES_ROOT_FOLDER + "/fortyfive_to_sixtyfour/"
	MIDDLEADULTHOOD_STAGE_ROUTE  = public_route + "/fortyfive-sixtyfour/"
	LATEADULTHOOD_STAGE_FOLDER   = STAGES_ROOT_FOLDER + "/sixtyfive_plus/"
	LATEADULTHOOD_STAGE_ROUTE    = public_route + "/sixtyfiveplus/"

	/**** PRIVATE CONSTANTS ****/
	PRIVATE_DELIVERY_FOLDER = PRIVATE_FOLDER + "/birth/delivery/"
	DELIVERY_ROUTE          = private_route + "/delivery/"

	/**** COOKIE VALUES ****/
	COOKIE_SECURE   bool = false
	COOKIE_HTTPONLY      = false
)

// Used to map the endpoint in a request to the matching folder on the filesystem. Some are the same but some are not
var ENDPOINT_TO_FOLDER = map[string]string{
	"ultrasounds":         "ultrasounds",
	"growingmommy":        "growing_mommy",
	"shower_reveal":       "shower_reveal",
	"dob":                 "day_of_birth",
	"delivery":            "delivery",
	"zero-one":            "zero_to_one",
	"one-two":             "one_to_two",
	"three-six":           "three_to_six",
	"seven-ten":           "seven_to_ten",
	"eleven-nineteen":     "eleven_to_nineteen",
	"twenty-fortyfour":    "twenty_to_fortyfour",
	"fortyfive-sixtyfour": "fortyfive_to_sixtyfour",
	"sixtyfiveplus":       "sixtyfive_plus",
}

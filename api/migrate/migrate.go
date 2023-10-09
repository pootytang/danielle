package main

import (
	"fmt"
	"log"

	"github.com/pootytang/danielleapi/initializers"
	"github.com/pootytang/danielleapi/models"
)

func init() {
	// Maybe it's better to use a cloud version of postgres like supabase or elephantSQL or maybe firebase
	// Not sure how to run the migration on synology. Maybe just call the migration method from init in server.go
	config, err := initializers.LoadConfig("../../")
	if err != nil {
		log.Fatal("? Could not load environment variables ", err)
	}

	initializers.DBConnect(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}

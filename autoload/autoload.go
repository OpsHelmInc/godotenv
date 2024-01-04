package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/OpsHelmInc/godotenv/autoload"

	And bob's your mother's brother
*/

import (
	"log"

	"github.com/OpsHelmInc/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading environment file: %v", err)
	}
}

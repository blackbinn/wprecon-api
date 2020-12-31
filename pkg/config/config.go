package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGOURI string
	APIPORT  int
)

func Loading() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APIPORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		APIPORT = 9000
	}

	MONGOURI = fmt.Sprintf("mongodb+srv://%s:%s@%s.9frcq.mongodb.net/%s?retryWrites=true&w=majority",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_DB"))
}

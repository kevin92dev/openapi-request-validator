package main

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	app "main/cmd/my_application"
	"os"
)

func init() {
	environment := os.Getenv("ENV")

	if "dev" == environment {
		log.SetOutput(os.Stdout)
	}
}

//go:embed api/docs/openapi.yml
var openapiDocs []byte

func main() {
	err := os.Setenv("OPENAPI_DOCS", string(openapiDocs))
	if err != nil {
		return
	}
	app.Start()
}

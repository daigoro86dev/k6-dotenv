package xk6dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/dotenv", new(Dotenv))
}

type Dotenv struct {
	EnvVar string
}

func (d *Dotenv) LoadEnvFileFromEnv() {
	switch env := os.Getenv("ENV"); env {
	case "PROD":
		loadFile("prod.env")
	case "DEV":
		loadFile("dev.env")
	}
}

func (d *Dotenv) GetEnvVar(varName string) string {
	d.EnvVar = os.Getenv(varName)
	return d.EnvVar
}

func loadFile(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatalf("Error loading %s file", fileName)
	}
}

package bootstrap

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/trillyai/backend-microservices/core/database/postgres"
)

var (
	envFiles = [3]string{".env.local", ".env.test", ".env.production"}
	Configs  = make(map[string]string)
)

func SetUpEnvironment() {
	loadEnvFileVars()
	loadFlagVars()
	loadOSVars()

	//TODO: validate configs here
	requiredEnvVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_DBNAME",

		"HTTP_PORT",
	}

	// Check if required environment variables exist
	for _, key := range requiredEnvVars {
		if _, exists := Configs[key]; !exists {
			log.Panicf("Missing required environment variable: %s", key)
		}
	}

	// Set DB parameters
	postgres.Host = Configs["DB_HOST"]
	postgres.User = Configs["DB_USER"]
	postgres.Password = Configs["DB_PASSWORD"]
	postgres.Dbname = Configs["DB_DBNAME"]
	port, err := strconv.Atoi(Configs["DB_PORT"])
	if err != nil {
		log.Panicf("invalid required environment variable for db port: %s", err)
	}
	postgres.Port = port
}

// loadEnvFileVars loads environment variables from dotenv files.
func loadEnvFileVars() {
	for _, fname := range envFiles {
		if file, err := os.Open(fname); err != nil {
		} else {
			if mapin, err := godotenv.Parse(file); err != nil {
				log.Printf("Error parsing dotenv file %s: %s", fname, err)
			} else {
				updateConfig(mapin)
				log.Printf("Environment variables loaded from file: %s", fname)
			}
		}
	}
}

// loadFlagVars loads environment variables from command-line flags.
func loadFlagVars() {
	args := os.Args[1:]
	keyValuePairs := make(map[string]string)
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) == 2 {
			keyValuePairs[parts[0]] = parts[1]
		}
	}
	updateConfig(keyValuePairs)
	log.Print("Environment variables loaded from command-line flags")
}

// loadOSVars loads environment variables from the operating system.
func loadOSVars() {
	envVariables := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			envVariables[pair[0]] = pair[1]
		}
	}
	updateConfig(envVariables)
	log.Print("Environment variables loaded from the operating system")
}

// updateConfig updates the configuration with new values.
func updateConfig(newConfig map[string]string) {
	for key, value := range newConfig {
		Configs[key] = value
	}
}

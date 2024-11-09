package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// AppConfig struct represents the application configuration.
type AppConfig struct {
	Port     int    `json:"port" env:"PORT"`
	DBType   string `json:"db_type" env:"DB_TYPE"`
	DBHost   string `json:"db_host" env:"DB_HOST"`
	DBPort   int    `json:"db_port" env:"DB_PORT"`
	DBUser   string `json:"db_user" env:"DB_USER"`
	DBPass   string `json:"db_pass" env:"DB_PASS"`
	DBName   string `json:"db_name" env:"DB_NAME"`
	AppName  string `json:"app_name" env:"APP_NAME"`
	LogLevel string `json:"log_level" env:"LOG_LEVEL"`
}

// NewConfig function creates a new AppConfig instance.
func NewConfig() (*AppConfig, error) {
	var config AppConfig

	// Initialize viper to use multiple configuration sources
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	// Check if the json file exists and load it first
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

		// Unmarshal the json config into the AppConfig struct
		if err := json.Unmarshal(viper.GetBytes(), &config); err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %w", err)
		}

	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Set the default values for any missing fields
	if config.Port == 0 {
		config.Port = 8080
	}
	if config.DBType == "" {
		config.DBType = "sqlite3"
	}
	if config.DBHost == "" {
		config.DBHost = "localhost"
	}
	if config.DBPort == 0 {
		config.DBPort = 5432
	}
	if config.DBUser == "" {
		config.DBUser = "postgres"
	}
	if config.DBName == "" {
		config.DBName = "mydb"
	}
	if config.AppName == "" {
		config.AppName = "myapp"
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}

	// Override the default values with environment variables if they are set
	viper.AutomaticEnv()
	if port := os.Getenv("PORT"); port != "" {
		config.Port, _ = strconv.Atoi(port)
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort, _ = strconv.Atoi(dbPort)
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	}
	if dbPass := os.Getenv("DB_PASS"); dbPass != "" {
		config.DBPass = dbPass
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	}
	if appName := os.Getenv("APP_NAME"); appName != "" {
		config.AppName = appName
	}
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		config.LogLevel = logLevel
	}

	return &config, nil
}

// GetConfig function returns the AppConfig instance.
func GetConfig() (*AppConfig, error) {
	if config == nil {
		var err error
		config, err = NewConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to create config: %w", err)
		}
	}

	return config, nil
}

// PrintConfig function prints the AppConfig instance.
func PrintConfig() error {
	config, err := GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	fmt.Printf("AppConfig:\n")
	fmt.Printf("  Port: %d\n", config.Port)
	fmt.Printf("  DBHost: %s\n", config.DBHost)
	fmt.Printf("  DBPort: %d\n", config.DBPort)
	fmt.Printf("  DBUser: %s\n", config.DBUser)
	fmt.Printf("  DBPass: %s\n", config.DBPass)
	fmt.Printf("  DBName: %s\n", config.DBName)
	fmt.Printf("  AppName: %s\n", config.AppName)
	fmt.Printf("  LogLevel: %s\n", config.LogLevel)

	return nil
}

// config instance
var config *AppConfig

func init() {
	// Print the config when the package is initialized
	if err := PrintConfig(); err != nil {
		log.Fatalf("Failed to print config: %v", err)
	}
}

func main() {
	flag.Parse()

	// Check if help flag was provided
	if *helpFlag {
		fmt.Println("Usage of " + os.Args[0] + ":")
		flag.Usage()
		os.Exit(0)
	}

	// Get the config and print it out
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}
	fmt.Printf("AppConfig:\n")
	fmt.Printf("  Port: %d\n", config.Port)
	fmt.Printf("  DBHost: %s\n", config.DBHost)
	fmt.Printf("  DBPort: %d\n", config.DBPort)
	fmt.Printf("  DBUser: %s\n", config.DBUser)
	fmt.Printf("  DBPass: %s\n", config.DBPass)
	fmt.Printf("  DBName: %s\n", config.DBName)
	fmt.Printf("  AppName: %s\n", config.AppName)
	fmt.Printf("  LogLevel: %s\n", config.LogLevel)

}

package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Type string `json:"db_type" mapstructure:"type"`
	Host string `json:"db_host" mapstructure:"host"`
	Port int    `json:"db_port" mapstructure:"port"`
	User string `json:"db_user" mapstructure:"user"`
	Pass string `json:"db_pass" mapstructure:"pass"`
	Name string `json:"db_name" mapstructure:"name"`
}

type AppConfig struct {
	DB       DBConfig `mapstructure:"db"`
	Port     int      `json:"port" mapstructure:"port"`
	LogLevel int      `json:"log_level" mapstructure:"log_level"`
}

func CobraSetup(cmd *cobra.Command) {
	// Default values

	// Load default values
	// Load config file
	// Load parameters
	// Load env vars

	//cmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Path to a config file")
	cmd.Flags().IntVar(&config.LogLevel, "log-level", int(slog.LevelInfo), "Log level (e.g. debug, info, warn, error)")
	cmd.Flags().IntVar(&config.Port, "port", config.Port, "Port to run the application on")
	cmd.Flags().StringVar(&config.DB.Type, "db-type", config.DB.Type, "Database type (e.g. postgres, mysql)")
	cmd.Flags().StringVar(&config.DB.Host, "db-host", config.DB.Host, "Database host")
	cmd.Flags().IntVar(&config.DB.Port, "db-port", config.DB.Port, "Database port")
	cmd.Flags().StringVar(&config.DB.User, "db-user", config.DB.User, "Database user")
	cmd.Flags().StringVar(&config.DB.Pass, "db-pass", config.DB.Pass, "Database password")
	cmd.Flags().StringVar(&config.DB.Name, "db-name", config.DB.Name, "Database name")

	SetLoggingLevel(config.LogLevel)
}

func SetLoggingLevel(level int) {
	// Set logging level
	lvl := new(slog.LevelVar)
	lvl.Set(slog.Level(level))
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: lvl})
	slog.SetDefault(slog.New(handler))
	slog.Info(fmt.Sprintf("Set log level to %d", level))
}

// LoadConfig function creates a new AppConfig instance.
func LoadConfig() (*AppConfig, error) {
	//cmd *cobra.Command, args []string
	config = &AppConfig{}

	// Initialize viper to use multiple configuration sources
	viper.SetEnvPrefix("dnh")
	viper.SetConfigFile(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	// Check if the json file exists and load it first
	if err := viper.ReadInConfig(); err == nil {
		slog.Info(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		slog.Warn(fmt.Sprintf("Failed to read config file: %s", err.Error()))
	}

	// Set the default values for any missing fields
	viper.SetDefault("port", 8088)
	viper.SetDefault("log_level", new(slog.LevelVar)) // Info by default
	viper.SetDefault("db_type", "sqlite3")
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 5432)
	viper.SetDefault("db_user", "postrges")
	viper.SetDefault("db_name", "dnh")

	// Override the default values with environment variables if they are set
	err := viper.Unmarshal(config)
	if err != nil {
		slog.Error("")
		return config, fmt.Errorf("Unable to decode into config struct, %v", err)
	}
	return config, nil
}

// GetConfig function returns the AppConfig instance.
func GetConfig() (*AppConfig, error) {
	if config == nil {
		var err error
		config, err = LoadConfig()
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
	fmt.Printf("  DBHost: %s\n", config.DB.Host)
	fmt.Printf("  DBPort: %d\n", config.DB.Port)
	fmt.Printf("  DBUser: %s\n", config.DB.User)
	fmt.Printf("  DBPass: %s\n", config.DB.Pass)
	fmt.Printf("  DBName: %s\n", config.DB.Name)
	fmt.Printf("  LogLevel: %d\n", config.LogLevel)

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

package config

import (
	"fmt"
	"strings"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds all application configuration
// This struct demonstrates how to structure configuration for dependency injection
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	App      AppConfig      `mapstructure:"app"`
}

// DatabaseConfig holds PostgreSQL configuration
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	SSLMode         string `mapstructure:"ssl_mode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// RabbitMQConfig holds RabbitMQ configuration
type RabbitMQConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	QueueName string `mapstructure:"queue_name"`
	Exchange  string `mapstructure:"exchange"`
}

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	Level   string `mapstructure:"level"`
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	NoColor bool   `mapstructure:"no_color"`
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

// NewConfig creates a new configuration instance using viper
// This demonstrates configuration management with the samber/do library
func NewConfig(i do.Injector) (*Config, error) {
	// Enable environment variable support
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	// Unmarshal configuration into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

// SetCobraFlags adds command line flags to the cobra command
// This method demonstrates how services can provide functionality through DI
func (cs *Config) SetCobraFlags(cmd *cobra.Command) {
	// Database flags
	cmd.PersistentFlags().String("database.host", "localhost", "Database host")
	cmd.PersistentFlags().Int("database.port", 5432, "Database port")
	cmd.PersistentFlags().String("database.user", "postgres", "Database user")
	cmd.PersistentFlags().String("database.password", "postgres", "Database password")
	cmd.PersistentFlags().String("database.database", "do_template_worker", "Database name")
	cmd.PersistentFlags().String("database.ssl_mode", "disable", "Database SSL mode")
	cmd.PersistentFlags().Int("database.max_open_conns", 25, "Database max open connections")
	cmd.PersistentFlags().Int("database.max_idle_conns", 25, "Database max idle connections")
	cmd.PersistentFlags().Int("database.conn_max_lifetime", 300, "Database connection max lifetime in seconds")

	// RabbitMQ flags
	cmd.PersistentFlags().String("rabbitmq.host", "localhost", "RabbitMQ host")
	cmd.PersistentFlags().Int("rabbitmq.port", 5672, "RabbitMQ port")
	cmd.PersistentFlags().String("rabbitmq.user", "guest", "RabbitMQ user")
	cmd.PersistentFlags().String("rabbitmq.password", "guest", "RabbitMQ password")
	cmd.PersistentFlags().String("rabbitmq.queue_name", "worker_queue", "RabbitMQ queue name")
	cmd.PersistentFlags().String("rabbitmq.exchange", "worker_exchange", "RabbitMQ exchange name")

	// Logger flags
	cmd.PersistentFlags().String("logger.level", "info", "Log level")
	cmd.PersistentFlags().String("logger.format", "console", "Log format")
	cmd.PersistentFlags().String("logger.output", "stdout", "Log output")
	cmd.PersistentFlags().Bool("logger.no_color", false, "Disable colored output")

	// App flags
	cmd.PersistentFlags().String("app.name", "do-template-worker", "Application name")
	cmd.PersistentFlags().String("app.version", "1.0.0", "Application version")
	cmd.PersistentFlags().String("app.environment", "development", "Application environment")
	cmd.PersistentFlags().Bool("app.debug", false, "Debug mode")

	// Bind all flags to viper for automatic configuration
	cs.bindFlagsToViper(cmd)
}

// bindFlagsToViper binds all cobra flags to viper
func (cs *Config) bindFlagsToViper(cmd *cobra.Command) {
	// Database flags
	viper.BindPFlag("database.host", cmd.PersistentFlags().Lookup("database.host"))
	viper.BindPFlag("database.port", cmd.PersistentFlags().Lookup("database.port"))
	viper.BindPFlag("database.user", cmd.PersistentFlags().Lookup("database.user"))
	viper.BindPFlag("database.password", cmd.PersistentFlags().Lookup("database.password"))
	viper.BindPFlag("database.database", cmd.PersistentFlags().Lookup("database.database"))
	viper.BindPFlag("database.ssl_mode", cmd.PersistentFlags().Lookup("database.ssl_mode"))
	viper.BindPFlag("database.max_open_conns", cmd.PersistentFlags().Lookup("database.max_open_conns"))
	viper.BindPFlag("database.max_idle_conns", cmd.PersistentFlags().Lookup("database.max_idle_conns"))
	viper.BindPFlag("database.conn_max_lifetime", cmd.PersistentFlags().Lookup("database.conn_max_lifetime"))

	// RabbitMQ flags
	viper.BindPFlag("rabbitmq.host", cmd.PersistentFlags().Lookup("rabbitmq.host"))
	viper.BindPFlag("rabbitmq.port", cmd.PersistentFlags().Lookup("rabbitmq.port"))
	viper.BindPFlag("rabbitmq.user", cmd.PersistentFlags().Lookup("rabbitmq.user"))
	viper.BindPFlag("rabbitmq.password", cmd.PersistentFlags().Lookup("rabbitmq.password"))
	viper.BindPFlag("rabbitmq.queue_name", cmd.PersistentFlags().Lookup("rabbitmq.queue_name"))
	viper.BindPFlag("rabbitmq.exchange", cmd.PersistentFlags().Lookup("rabbitmq.exchange"))

	// Logger flags
	viper.BindPFlag("logger.level", cmd.PersistentFlags().Lookup("logger.level"))
	viper.BindPFlag("logger.format", cmd.PersistentFlags().Lookup("logger.format"))
	viper.BindPFlag("logger.output", cmd.PersistentFlags().Lookup("logger.output"))
	viper.BindPFlag("logger.no_color", cmd.PersistentFlags().Lookup("logger.no_color"))

	// App flags
	viper.BindPFlag("app.name", cmd.PersistentFlags().Lookup("app.name"))
	viper.BindPFlag("app.version", cmd.PersistentFlags().Lookup("app.version"))
	viper.BindPFlag("app.environment", cmd.PersistentFlags().Lookup("app.environment"))
	viper.BindPFlag("app.debug", cmd.PersistentFlags().Lookup("app.debug"))
}

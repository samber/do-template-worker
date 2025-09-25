package cli

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do-template-worker/pkg/workers"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

// CLI represents the command line interface service
// This demonstrates how to create a CLI service with dependency injection.
type CLI struct {
	config      *config.Config `do:""`
	injector    do.Injector    `do:""`
	rootCommand *cobra.Command
}

// NewCLI creates a new CLI service with dependency injection support.
func NewCLI(i do.Injector) (*CLI, error) {
	cli := do.MustInvokeStruct[*CLI](i)

	// Create the root command
	cli.rootCommand = &cobra.Command{
		Use:     cli.config.App.Name,
		Short:   "A template worker application using samber/do dependency injection",
		Long:    "A comprehensive template project demonstrating the github.com/samber/do dependency injection library with PostgreSQL and RabbitMQ integration",
		Version: cli.config.App.Version,
	}

	// Add persistent flags using dependency injection
	cli.setupPersistentFlags()

	// Add commands
	cli.setupCommands()

	return cli, nil
}

// setupPersistentFlags adds global flags to the CLI.
func (cli *CLI) setupPersistentFlags() {
	// Use the config service to set up all configuration flags
	// This demonstrates dependency injection for configuration management
	cli.config.SetCobraFlags(cli.rootCommand)
}

// setupCommands adds subcommands to the CLI.
func (cli *CLI) setupCommands() {
	// Add producer command
	cli.rootCommand.AddCommand(cli.newProducerCommand())

	// Add consumer command
	cli.rootCommand.AddCommand(cli.newConsumerCommand())

	// Add serve command
	cli.rootCommand.AddCommand(cli.newServeCommand())

	// Add migrate command
	cli.rootCommand.AddCommand(cli.newMigrateCommand())

	// Add health command
	cli.rootCommand.AddCommand(cli.newHealthCommand())

	// Add version command
	cli.rootCommand.AddCommand(cli.newVersionCommand())
}

// newProducerCommand creates the producer command.
func (cli *CLI) newProducerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "producer",
		Short: "Start the producer worker",
		Long:  "Start the producer worker that creates messages periodically",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting producer worker...")
			cli.runProducer()
		},
	}
}

// newConsumerCommand creates the consumer command.
func (cli *CLI) newConsumerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "Start the consumer worker",
		Long:  "Start the consumer worker that processes messages and calls UserRepository",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting consumer worker...")
			cli.runConsumer()
		},
	}
}

// newServeCommand creates the serve command.
func (cli *CLI) newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the worker service",
		Long:  "Start the do-template-worker service with dependency injection",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting worker service...")
			// This will be implemented to use the dependency injection container
		},
	}
}

// newMigrateCommand creates the migrate command.
func (cli *CLI) newMigrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Long:  "Run database migrations using the configured database connection",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running database migrations...")
			// This will be implemented to use the dependency injection container
		},
	}
}

// newHealthCommand creates the health command.
func (cli *CLI) newHealthCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Check service health",
		Long:  "Check the health of all services and dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking service health...")
			// This will be implemented to use the dependency injection container
		},
	}
}

// newVersionCommand creates the version command.
func (cli *CLI) newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Show detailed version and build information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s\n", cli.config.App.Name, cli.config.App.Version)
		},
	}
}

// RootCommand returns the root cobra command.
func (cli *CLI) RootCommand() *cobra.Command {
	return cli.rootCommand
}

// Execute executes the CLI with the given arguments.
func (cli *CLI) Execute() error {
	return cli.rootCommand.Execute()
}

// AddCommand adds a new command to the CLI.
func (cli *CLI) AddCommand(command *cobra.Command) {
	cli.rootCommand.AddCommand(command)
}

// runProducer starts the producer worker with graceful shutdown
// This method demonstrates how to run a worker with dependency injection and signal handling.
func (cli *CLI) runProducer() {
	// Get services from dependency injection container
	producerWorker := do.MustInvoke[*workers.ProducerWorker](cli.injector)
	logger := do.MustInvoke[zerolog.Logger](cli.injector)

	// Start the producer worker
	if err := producerWorker.Start(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start producer worker")
	}
}

// runConsumer starts the consumer worker with graceful shutdown
// This method demonstrates how to run a worker with dependency injection and signal handling.
func (cli *CLI) runConsumer() {
	// Get services from dependency injection container
	consumerWorker := do.MustInvoke[*workers.ConsumerWorker](cli.injector)
	logger := do.MustInvoke[zerolog.Logger](cli.injector)

	// Start the consumer worker
	if err := consumerWorker.Start(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start consumer worker")
	}
}

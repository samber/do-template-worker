package main

import (
	"github.com/rs/zerolog"
	"github.com/samber/do-template-worker/pkg"
	"github.com/samber/do-template-worker/pkg/cli"
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do-template-worker/pkg/repositories"
	"github.com/samber/do-template-worker/pkg/workers"
	"github.com/samber/do/v2"
)

func main() {
	// Initialize the dependency injection injector
	// This is the core component of the samber/do library that manages all services
	injector := do.New(
		pkg.BasePackage,
		repositories.Package,
		workers.WorkerPackage,
	)

	// Get services from dependency injection container
	appConfig := do.MustInvoke[*config.Config](injector)
	appLogger := do.MustInvoke[zerolog.Logger](injector)
	cliService := do.MustInvoke[*cli.CLI](injector)

	// Start the application
	appLogger.Info().Str("app_name", appConfig.App.Name).
		Str("version", appConfig.App.Version).
		Str("environment", appConfig.App.Environment).
		Msg("Starting do-template-worker application")

	// Execute the CLI - this will handle all command parsing and execution
	if err := cliService.Execute(); err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to execute CLI")
	}

	_, _ = injector.ShutdownOnSignals()
}

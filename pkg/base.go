package pkg

import (
	"github.com/samber/do-template-worker/pkg/cli"
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do-template-worker/pkg/logger"
	"github.com/samber/do-template-worker/pkg/repositories"
	"github.com/samber/do/v2"
)

var BasePackage = do.Package(
	do.Lazy(config.NewConfig),
	do.Lazy(cli.NewCLI),
	do.Lazy(logger.NewLogger),
	do.Lazy(repositories.NewDatabase),
	do.Lazy(repositories.NewUserRepository),
)

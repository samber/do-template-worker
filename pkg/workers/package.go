package workers

import (
	"github.com/samber/do-template-worker/pkg/rabbitmq"
	"github.com/samber/do/v2"
)

// WorkerPackage provides worker services to the dependency injector
// This package demonstrates how to organize worker services using the samber/do library.
var WorkerPackage = do.Package(
	do.Lazy(rabbitmq.ProvideRabbitMQConfig),
	do.Lazy(rabbitmq.NewRabbitMQService),
	do.Lazy(NewProducerWorker),
	do.Lazy(NewConsumerWorker),
)

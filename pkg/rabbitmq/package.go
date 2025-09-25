package rabbitmq

import (
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do/v2"
)

// ProvideRabbitMQConfig provides RabbitMQ configuration to the dependency injector
// This function demonstrates how to provide configuration using the samber/do library.
func ProvideRabbitMQConfig(injector do.Injector) (*Config, error) {
	appConfig := do.MustInvoke[*config.Config](injector)

	// Convert from config.RabbitMQConfig to rabbitmq.Config
	return &Config{
		Host:      appConfig.RabbitMQ.Host,
		Port:      appConfig.RabbitMQ.Port,
		User:      appConfig.RabbitMQ.User,
		Password:  appConfig.RabbitMQ.Password,
		QueueName: appConfig.RabbitMQ.QueueName,
		Exchange:  appConfig.RabbitMQ.Exchange,
	}, nil
}

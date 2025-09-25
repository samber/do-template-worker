package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do-template-worker/pkg/rabbitmq"
	"github.com/samber/do-template-worker/pkg/repositories"
	"github.com/samber/do/v2"
)

// ProducerWorker is a worker that produces messages to RabbitMQ
// This struct demonstrates how to implement a producer worker with dependency injection
type ProducerWorker struct {
	rabbitMQ *rabbitmq.RabbitMQService
	userRepo repositories.UserRepository
	logger   *zerolog.Logger
	config   *config.Config
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewProducerWorker creates a new producer worker instance
// This function demonstrates how to initialize a producer with dependency injection
func NewProducerWorker(injector do.Injector) (*ProducerWorker, error) {
	ctx, cancel := context.WithCancel(context.Background())

	return &ProducerWorker{
		rabbitMQ: do.MustInvoke[*rabbitmq.RabbitMQService](injector),
		userRepo: do.MustInvoke[repositories.UserRepository](injector),
		logger:   do.MustInvoke[*zerolog.Logger](injector),
		config:   do.MustInvoke[*config.Config](injector),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Start starts the producer worker
// This method demonstrates how to start a producer worker with dependency injection
func (w *ProducerWorker) Start() error {
	w.logger.Info().Msg("Starting producer worker")

	// Start producing messages periodically
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-w.ctx.Done():
				w.logger.Info().Msg("Producer worker stopped")
				return
			case <-ticker.C:
				if err := w.produceMessage(); err != nil {
					w.logger.Error().Err(err).Msg("Failed to produce message")
				}
			}
		}
	}()

	return nil
}

// Shutdown stops the producer worker
// This method demonstrates how to stop a producer worker with dependency injection
func (w *ProducerWorker) Shutdown() error {
	w.logger.Info().Msg("Stopping producer worker")
	w.cancel()
	return nil
}

// produceMessage produces a message to RabbitMQ
// This method demonstrates how to produce a message with dependency injection
func (w *ProducerWorker) produceMessage() error {
	// Create a message
	message := WorkerMessage{
		Action: "create_user",
		Payload: UserPayload{
			Name:  fmt.Sprintf("User_%d", time.Now().Unix()),
			Email: fmt.Sprintf("user_%d@example.com", time.Now().Unix()),
		},
		ID: fmt.Sprintf("msg_%d", time.Now().UnixNano()),
	}

	// Serialize message
	messageData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Publish message
	if err := w.rabbitMQ.PublishMessage(messageData); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	w.logger.Info().Str("message_id", message.ID).Msg("Produced message")
	return nil
}

package workers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"github.com/samber/do-template-worker/pkg/config"
	"github.com/samber/do-template-worker/pkg/rabbitmq"
	"github.com/samber/do-template-worker/pkg/repositories"
	"github.com/samber/do/v2"
)

// ConsumerWorker is a worker that consumes messages from RabbitMQ
// This struct demonstrates how to implement a consumer worker with dependency injection
type ConsumerWorker struct {
	rabbitMQ *rabbitmq.RabbitMQService
	userRepo repositories.UserRepository
	logger   *zerolog.Logger
	config   *config.Config
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewConsumerWorker creates a new consumer worker instance
// This function demonstrates how to initialize a consumer with dependency injection
func NewConsumerWorker(injector do.Injector) (*ConsumerWorker, error) {
	ctx, cancel := context.WithCancel(context.Background())

	return &ConsumerWorker{
		rabbitMQ: do.MustInvoke[*rabbitmq.RabbitMQService](injector),
		userRepo: do.MustInvoke[repositories.UserRepository](injector),
		logger:   do.MustInvoke[*zerolog.Logger](injector),
		config:   do.MustInvoke[*config.Config](injector),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Start starts the consumer worker
// This method demonstrates how to start a consumer worker with dependency injection
func (w *ConsumerWorker) Start() error {
	w.logger.Info().Msg("Starting consumer worker")

	// Start consuming messages
	go func() {
		// Create a new channel for each consumer instance
		msgChan, err := w.rabbitMQ.ConsumeMessage()
		if err != nil {
			w.logger.Error().Err(err).Msg("Failed to start consuming messages")
			return
		}

		for {
			select {
			case <-w.ctx.Done():
				w.logger.Info().Msg("Consumer worker stopped")
				return
			case msg, ok := <-msgChan:
				if !ok {
					w.logger.Info().Msg("Message channel closed")
					return
				}

				if err := w.processMessage(msg); err != nil {
					w.logger.Error().Err(err).Msg("Failed to process message")
					msg.Nack(false, true)
				} else {
					msg.Ack(false)
				}
			}
		}
	}()

	return nil
}

// Shutdown stops the consumer worker
// This method demonstrates how to stop a consumer worker with dependency injection
func (w *ConsumerWorker) Shutdown() error {
	w.logger.Info().Msg("Stopping consumer worker")
	w.cancel()
	return nil
}

// processMessage processes a message from RabbitMQ
// This method demonstrates how to process a message with dependency injection and UserRepository
func (w *ConsumerWorker) processMessage(msg amqp091.Delivery) error {
	// Deserialize message
	var message WorkerMessage
	if err := json.Unmarshal(msg.Body, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	w.logger.Info().
		Str("message_id", message.ID).
		Str("action", message.Action).
		Msg("Processing message")

	// Process message based on action
	switch message.Action {
	case "create_user":
		return w.handleCreateUser(message.Payload)
	default:
		w.logger.Warn().Str("action", message.Action).Msg("Unknown action")
		return nil
	}
}

// handleCreateUser handles the create user action
// This method demonstrates how to use UserRepository with dependency injection
func (w *ConsumerWorker) handleCreateUser(payload interface{}) error {
	userPayload, ok := payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid payload type")
	}

	name, ok := userPayload["name"].(string)
	if !ok {
		return fmt.Errorf("name not found in payload")
	}

	email, ok := userPayload["email"].(string)
	if !ok {
		return fmt.Errorf("email not found in payload")
	}

	// Create user using UserRepository
	user := &repositories.User{
		Name:  name,
		Email: email,
	}

	createdUser, err := w.userRepo.CreateUser(w.ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	w.logger.Info().
		Int64("user_id", createdUser.ID).
		Str("user_name", createdUser.Name).
		Str("user_email", createdUser.Email).
		Msg("Created user from message")

	return nil
}

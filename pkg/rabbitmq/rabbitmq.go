package rabbitmq

import (
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/samber/do/v2"
)

// RabbitMQService represents a RabbitMQ connection and channel manager
// This struct demonstrates how to manage RabbitMQ connections with dependency injection
type RabbitMQService struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	config  *Config `do:""`
}

// Config holds RabbitMQ configuration
type Config struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	QueueName string `mapstructure:"queue_name"`
	Exchange  string `mapstructure:"exchange"`
}

// NewRabbitMQService creates a new RabbitMQ service instance
// This function demonstrates how to initialize a message broker service with dependency injection
func NewRabbitMQService(injector do.Injector) (*RabbitMQService, error) {
	// Get configuration from injector
	config := do.MustInvoke[*Config](injector)

	// Build connection URL
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)

	// Connect to RabbitMQ
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Create channel
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create RabbitMQ channel: %w", err)
	}

	// Declare exchange
	err = channel.ExchangeDeclare(
		config.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare queue
	_, err = channel.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange
	err = channel.QueueBind(
		config.QueueName,
		config.QueueName,
		config.Exchange,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue to exchange: %w", err)
	}

	return &RabbitMQService{
		conn:    conn,
		channel: channel,
		config:  config,
	}, nil
}

// PublishMessage publishes a message to the RabbitMQ queue
// This method demonstrates how to send messages using dependency injection
func (r *RabbitMQService) PublishMessage(message []byte) error {
	return r.channel.Publish(
		r.config.Exchange,
		r.config.QueueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
			Timestamp:   time.Now(),
		},
	)
}

// ConsumeMessage starts consuming messages from the RabbitMQ queue
// This method demonstrates how to consume messages using dependency injection
func (r *RabbitMQService) ConsumeMessage() (<-chan amqp091.Delivery, error) {
	return r.channel.Consume(
		r.config.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

// Close closes the RabbitMQ connection and channel
// This method demonstrates proper resource cleanup in dependency injection
func (r *RabbitMQService) Shutdown() error {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	return nil
}

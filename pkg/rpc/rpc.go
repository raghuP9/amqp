package rpc

import (
	"context"

	"github.com/raghuP9/amqp/pkg/rpc/rmq"
	"github.com/streadway/amqp"
)

// RabbitMQRPC ...
type RabbitMQRPC interface {
	ExchangeDeclare(string, *rmq.DeclareExchangeOpts) error
	QueueDeclare(string, *rmq.DeclareQueueOpts) (amqp.Queue, error)
	QueueBind(string, string, string, *rmq.QueueBindOpts) error
	Publish(amqp.Publishing, string, string, *rmq.PublishOpts) error
	Subscribe(
		context.Context,
		string,
		string,
		bool,
		bool,
		func(amqp.Delivery) (amqp.Publishing, error),
	) error
}
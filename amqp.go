package kontrakto

import amqp "github.com/taranovegor/pkg.amqp"

const (
	AmqpShortLinkValidateToken    = "short_link.validate.token"
	AmqpShortLinkValidateLocation = "short_link.validate.location"
	AmqpShortLinkWriteCreate      = "short_link.write.create"
	AmqpShortLinkWriteUpdate      = "short_link.write.update"
	AmqpShortLinkWriteDelete      = "short_link.write.delete"
	AmqpShortLinkReadGet          = "short_link.read.get"
	AmqpShortLinkReadPaginate     = "short_link.read.paginate"
)

func AmqpConfig() amqp.Config {
	return amqp.NewConfig(
		map[string]amqp.ConsumerConfig{
			AmqpShortLinkValidateToken:    {Queue: AmqpShortLinkValidateToken, Exclusive: true, NoWait: true},
			AmqpShortLinkValidateLocation: {Queue: AmqpShortLinkValidateLocation, Exclusive: true, NoWait: true},
			AmqpShortLinkWriteCreate:      {Queue: AmqpShortLinkWriteCreate, Exclusive: true, NoWait: true},
			AmqpShortLinkWriteUpdate:      {Queue: AmqpShortLinkWriteUpdate, Exclusive: true, NoWait: true},
			AmqpShortLinkWriteDelete:      {Queue: AmqpShortLinkWriteDelete, Exclusive: true, NoWait: true},
			AmqpShortLinkReadGet:          {Queue: AmqpShortLinkReadGet, Exclusive: true, NoWait: true},
			AmqpShortLinkReadPaginate:     {Queue: AmqpShortLinkReadPaginate, Exclusive: true, NoWait: true},
		},
		map[string]amqp.ExchangeConfig{},
		map[string]amqp.QueueConfig{
			AmqpShortLinkValidateToken:    {AutoDelete: true},
			AmqpShortLinkValidateLocation: {AutoDelete: true},
			AmqpShortLinkWriteCreate:      {Durable: true},
			AmqpShortLinkWriteUpdate:      {Durable: true},
			AmqpShortLinkWriteDelete:      {Durable: true},
			AmqpShortLinkReadGet:          {AutoDelete: true},
			AmqpShortLinkReadPaginate:     {AutoDelete: true},
		},
		map[string][]amqp.QueueBindConfig{},
		map[string]amqp.ProducerConfig{
			AmqpShortLinkValidateToken:    {Key: AmqpShortLinkValidateToken},
			AmqpShortLinkValidateLocation: {Key: AmqpShortLinkValidateLocation},
			AmqpShortLinkWriteCreate:      {Key: AmqpShortLinkWriteCreate},
			AmqpShortLinkWriteUpdate:      {Key: AmqpShortLinkWriteUpdate},
			AmqpShortLinkWriteDelete:      {Key: AmqpShortLinkWriteDelete},
			AmqpShortLinkReadGet:          {Key: AmqpShortLinkReadGet},
			AmqpShortLinkReadPaginate:     {Key: AmqpShortLinkReadPaginate},
		},
		map[interface{}]amqp.RouteConfig{
			ValidateShortLinkToken{}:    {Producer: AmqpShortLinkValidateToken},
			ValidateShortLinkLocation{}: {Producer: AmqpShortLinkValidateLocation},
			CreateShortLink{}:           {Producer: AmqpShortLinkWriteCreate},
			UpdateShortLink{}:           {Producer: AmqpShortLinkWriteUpdate},
			DeleteShortLink{}:           {Producer: AmqpShortLinkWriteDelete},
			GetShortLink{}:              {Producer: AmqpShortLinkReadGet},
			PaginateShortLinks{}:        {Producer: AmqpShortLinkReadPaginate},
		},
	)
}

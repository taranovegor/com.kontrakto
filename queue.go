package kontrakto

const (
	QueueLigilo       = "ligilo"
	QueueNotification = "notification"
)

type QueueProducers map[interface{}][]string

type QueueConfig struct {
	Producers QueueProducers
}

func GetQueueConfig() QueueConfig {
	return QueueConfig{
		Producers: QueueProducers{
			CreateShortLink{}: {QueueLigilo},
			UpdateShortLink{}: {QueueLigilo},
			DeleteShortLink{}: {QueueLigilo},
		},
	}
}

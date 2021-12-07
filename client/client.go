package client

type Client interface {
	Subscribe(topic string, qos byte, callback MessageHandler) string
}

type MessageHandler func(string2 string)

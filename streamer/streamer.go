package streamer

import (
	"Go_mock_try/client"
)

type streamer struct {
	client client.Client
}

func NewStreamer(client client.Client) streamer {
	return streamer{client: client}
}

func (s *streamer) CallSubcribe() string {

	var method client.MessageHandler = func(string2 string) {
		print(string2)
	}

	return s.client.Subscribe("topic1111", byte(1), method)
}

func (s *streamer) RandFunc() string {
	return "random function called "
}

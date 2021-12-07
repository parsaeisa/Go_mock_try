package test

import (
	"testing"

	"Go_mock_try/mocks"
	"Go_mock_try/streamer"
	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockclient(ctrl)
	s := streamer.NewStreamer(m)

	//var method client.MessageHandler = func(string2 string) {
	//	print(string2)
	//}

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	m.
		EXPECT().
		Subscribe("topic", byte(1), gomock.Any()).MaxTimes(0).
		//DoAndReturn(func(_ int) int {
		//	time.Sleep(1*time.Second)
		//	return 101
		//}).
		AnyTimes()

	s.RandFunc()
}

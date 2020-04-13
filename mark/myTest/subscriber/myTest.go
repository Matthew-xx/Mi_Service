package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	myTest "mark/myTest/proto/myTest"
)

type MyTest struct{}

func (e *MyTest) Handle(ctx context.Context, msg *myTest.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *myTest.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

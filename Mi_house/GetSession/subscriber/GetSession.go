package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	GetSession "Mi_house/GetSession/proto/GetSession"
)

type GetSession struct{}

func (e *GetSession) Handle(ctx context.Context, msg *GetSession.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GetSession.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

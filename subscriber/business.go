package subscriber

import (
	"context"
	"github.com/micro/go-log"

	business "github.com/rockdragon/micro_example/proto/business"
)

type Api struct{}

func (e *Api) Handle(ctx context.Context, msg *business.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *business.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

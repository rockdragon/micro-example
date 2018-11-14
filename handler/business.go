package handler

import (
	"context"

	"github.com/micro/go-log"

	business "github.com/rockdragon/micro_example/proto/business"
)

type Api struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Api) Call(ctx context.Context, req *business.Request, rsp *business.Response) error {
	log.Log("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Api) Stream(ctx context.Context, req *business.StreamingRequest, stream business.API_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&business.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Api) PingPong(ctx context.Context, stream business.API_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&business.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

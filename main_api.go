package main

import (
	"encoding/json"
	"github.com/rockdragon/micro_example/utils"
	"log"
	"strings"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	example "github.com/rockdragon/micro_example/proto/example"

	"context"
)

type Example struct {
	Client example.ExampleService
}

func (s *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received API request from: micro_example/example/call")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest(utils.ApiName, "Name cannot be blank")
	}

	response, err := s.Client.Call(ctx, &example.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		log.Printf("Error occured: %v", err)
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name(utils.ApiName),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Example{Client: example.NewExampleService(utils.SrvName, service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

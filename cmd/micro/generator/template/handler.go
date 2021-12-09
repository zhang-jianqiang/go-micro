package template

// HandlerFNC is the handler template used for new function projects.
var HandlerFNC = `package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type {{title .Service}} struct{}

func (e *{{title .Service}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

// HandlerSRV is the handler template used for new service projects.
var HandlerSRV = `package handler

import (
	"context"
	"io"
	"time"
    "{{.Vendor}}{{.Service}}/domain/service"

	log "go-micro.dev/v4/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type {{title .Service}} struct{
	{{title .Service}}DataService service.I{{title .Service}}DataService
}

func (e *{{title .Service}}) FindOneById(ctx context.Context, req *pb.FindOneByIdRequest, rsp *pb.FindOneByIdRequestResponse) error {
	log.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

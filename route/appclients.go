package route

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"time"
)

func (p *resourcesPool) TryAddClient(address string) error {
	if _, ok := p.clients[address]; ok {
		return ErrClientExists
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	appCli := NewAppServeClient(conn)
	state := conn.GetState()
	if state == connectivity.Connecting {
		p.clients[address] = appCli
		return nil
	}
	return ErrGprcServerConnFailed
}

func (p *resourcesPool) RemoveClient(address string) {
	delete(p.clients, address)
}

func (p *resourcesPool) callAppAction(address string, request proto.Message) (*ServerComResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return p.clients[address].Request(ctx, request.(*ClientComRequest))
}

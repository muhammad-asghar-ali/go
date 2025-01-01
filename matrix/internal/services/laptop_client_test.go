package services_test

import (
	"context"
	"net"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"matrix/internal/fns"
	"matrix/internal/pb"
	"matrix/internal/pb/pbconnect"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	addr := _test_server(t)
	client := _new_client(addr)

	laptop := fns.NewLaptop()
	lpq := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	req := connect.NewRequest(lpq)

	res, err := client.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res.Msg)
	require.Equal(t, laptop.Id, res.Msg.GetId())
}

func _test_server(t *testing.T) string {
	server := grpc.NewServer()

	l, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go server.Serve(l)

	return l.Addr().String()
}

func _new_client(add string) pbconnect.LaptopServiceClient {
	c := &http.Client{}

	client := pbconnect.NewLaptopServiceClient(c, add)

	return client
}

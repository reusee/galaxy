package main

import (
	"net"

	"golang.org/x/net/context"

	"github.com/reusee/galaxy/protos/etc"
	"google.golang.org/grpc"
)

const (
	listenAddr = ":59876"
)

type server struct{}

func (s *server) Get(ctx context.Context, req *etc.GetReq) (reply *etc.GetReply, err error) {
	reply = &etc.GetReply{
		Value: "foo",
	}
	return
}

func main() {
	ln, err := net.Listen("tcp", listenAddr)
	ce(err, "listen")
	s := grpc.NewServer()
	etc.RegisterEtcServer(s, &server{})
	s.Serve(ln)
}

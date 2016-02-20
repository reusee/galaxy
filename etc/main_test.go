package main

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/reusee/galaxy/protos/etc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	ln, err := net.Listen("tcp", listenAddr)
	ce(err, "listen")
	s := grpc.NewServer()
	etc.RegisterEtcServer(s, &server{})
	go s.Serve(ln)
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	conn, err := grpc.Dial(listenAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()
	c := etc.NewEtcClient(conn)

	rep, err := c.Get(context.Background(), &etc.GetReq{})
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	fmt.Printf("%s\n", rep.Value)
}

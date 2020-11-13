package Server

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	istener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	msgpb.RegisterMessageServiceServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (m *MessageServiceServer) CreateBlog(ctx context.Context, req *msgpb.CreateInfoReq) (*msgpb.CreateInfoRes, error) {

	
}
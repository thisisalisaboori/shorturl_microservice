package main

import (
	"context"
	fmt "fmt"
	"net"
	app "shorturl/Applocation"
	pb "shorturl/services"

	"google.golang.org/grpc"
)

type Server struct{}

func (T *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.Response, error) {

	da := app.GetAppLicationService(0)
	result := da.SetLink(*req)
	return &result, nil
}
func (T *Server) Get(ctx context.Context, req *pb.GetLink) (*pb.Response, error) {
	//fmt.Println("aliiiiiiiiiiiii", req.Url)
	da := app.GetAppLicationService(0)

	r := da.GetLink(req.Url)
	result := pb.Response{Url: r, Shorturl: req.Url}
	//fmt.Printf("link=%s", r)
	return &result, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}
	var server Server
	srv := grpc.NewServer()
	pb.RegisterShortLinkServer(srv, &server)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

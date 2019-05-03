package main

import (
	"fmt"
	"net"

	"grpcdemo/pbclass"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Login(ctx context.Context, usr *pbclass.UserInfo) (*pbclass.FuncResponse, error) {
	fmt.Println("server Login() UserInfo:", usr)
	usr.Id = 100
	strId := "100"
	return &pbclass.FuncResponse{Reply: strId}, nil
}

func (s *server) Logout(ctx context.Context, uid *pbclass.UserID) (*pbclass.FuncResponse, error) {
	fmt.Println("server logout() UserID:", uid)
	return &pbclass.FuncResponse{Reply: "Logout Successed."}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Println("failed to listen:", err)
	}
	s := grpc.NewServer()
	pbclass.RegisterUserServiceServer(s, &server{})
	fmt.Println("success")
	s.Serve(lis)
}

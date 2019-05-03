package main

import (
	"fmt"
	"strconv"

	"grpcdemo/pbclass"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
	}
	defer conn.Close()
	c := pbclass.NewUserServiceClient(conn)

	var status pbclass.UserStatus
	status = pbclass.UserStatus_ONLINE

	userInfo := &pbclass.UserInfo{
		//Id:     10,         //proto.Int32(10),
		Name:   "XCL-gRPC", //proto.String("XCL-gRPC"),
		Status: status,
	}

	r, err := c.Login(context.Background(), userInfo)
	if err != nil {
		fmt.Println("登录失败! ", err)
	}
	fmt.Println("Login():", r)

	uid, err := strconv.ParseInt(r.Reply, 10, 32)
	if err != nil {
		fmt.Println("非数字 ", err)
	}

	userID := &pbclass.UserID{Id: int32(uid)}
	out, err := c.Logout(context.Background(), userID)
	fmt.Println("Logout():", out)
	a := ""
	fmt.Scan(&a)
	
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-example/auth"
	"time"
)

const port = 5001

func main() {
	//grpc 会使用 backoff 算法自动重连, 可以设置最大的重连时间
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBackoffMaxDelay(time.Second*3))

	if err != nil {
		panic(err)
	}

	go func() {
		//检测连接状态变化
		for {
			conn.WaitForStateChange(context.TODO(), conn.GetState())
			fmt.Println(conn.GetState())
		}
	}()

	client := rpc_auth.NewAuthClient(conn)

	resp, err := client.Login(context.TODO(), &rpc_auth.LoginRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Token)

	friends, err := client.GetUserInfo(context.TODO(), &rpc_auth.Token{Token: resp.Token})
	fmt.Println(friends)

	c := make(chan int)
	<-c
}

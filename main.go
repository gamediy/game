package main

import (
	"context"
	"fmt"
	"game/app/user/api/withdraw/withdraw"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// 创建与gRPC服务器的连接
	conn, err := grpc.Dial("10.10.46.4:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// 使用连接初始化HelloService客户端
	c := withdraw.NewWithdrawServiceClient(conn)

	// 设置一个超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 调用SayHello方法
	r, err := c.Submit(ctx, &withdraw.SubmitRequest{
		Amount: 1,
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s\n", r.Code)
}

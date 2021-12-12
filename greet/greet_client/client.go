package main

import (
	"context"
	"fmt"
	"go_grcp/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	fmt.Println("Hello, I'm a client")

	//withInscureはsslで通信するという意味
	//grcpをセキュアに行うオプション
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	//プログラムの最後にdeferでclientを閉める
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Println("Created client: %f", c)

	//doUnary(c)

	//ストリーミング用
	//doServerStreaming(c)

	doClientStreaming(c)
}

//Unaryでの実行結果
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "nekomamushi",
			LastName:  "aaaa"},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	//実行結果
	log.Printf("Response from Great: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "nekomamushi",
			LastName:  "timinekotimitimi",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet Server Streaming%v", err)
	}

	//ループでデータの受け取りを待っている

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//we're reached the end of the stream'
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
	}

	//リクエストは送らなくても良い
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet:%v", err)
	}

	//リクエストの送信を一つずつ行っている
	for _, req := range requests {
		fmt.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while Recv response form longGreetResponse:%v", err)
	}
	fmt.Printf("longGreet Response: %v\n", res)
}

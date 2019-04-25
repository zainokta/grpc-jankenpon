package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "grpc-jankenpon/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGameClient(conn)

	fmt.Print("Pilih Batu/Gunting/Kertas : \n")

	reader := bufio.NewReader(os.Stdin)
	request, _ := reader.ReadString('\n')

	ctx := context.Background()

	r, err := c.GameStart(ctx, &pb.ClientOption{Request: request})
	if err != nil {
		log.Fatalf("no response: %v", err)
	}
	log.Printf("Result: %s", r.Result)

	fmt.Scanln()
}

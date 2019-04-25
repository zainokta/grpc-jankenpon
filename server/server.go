package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "grpc-jankenpon/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type server struct{}

func (s *server) GameStart(ctx context.Context, in *pb.ClientOption) (*pb.Result, error) {
	fmt.Print("Pilih Batu/Gunting/Kertas : \n")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')

	answer := strings.TrimSuffix(choice, "\r\n")

	clientInput := strings.TrimSuffix(in.Request, "\r\n")

	// ini batu
	if strings.ToLower(clientInput) == "batu" && strings.ToLower(answer) == "gunting" {
		fmt.Printf("You Lose\n")
		return &pb.Result{Result: "You Win\n"}, nil
	} else if strings.ToLower(clientInput) == "batu" && strings.ToLower(answer) == "kertas" {
		fmt.Printf("You Win\n")
		return &pb.Result{Result: "You Lose\n"}, nil
	} else if strings.ToLower(clientInput) == "batu" && strings.ToLower(answer) == "batu" {
		fmt.Printf("Draw\n")
		return &pb.Result{Result: "Draw\n"}, nil
	}

	//ini gunting
	if strings.ToLower(clientInput) == "gunting" && strings.ToLower(answer) == "gunting" {
		fmt.Printf("Draw\n")
		return &pb.Result{Result: "Draw\n"}, nil
	} else if strings.ToLower(clientInput) == "gunting" && strings.ToLower(answer) == "kertas" {
		fmt.Printf("You Lose\n")
		return &pb.Result{Result: "You Win\n"}, nil
	} else if strings.ToLower(clientInput) == "gunting" && strings.ToLower(answer) == "batu" {
		fmt.Printf("You Win\n")
		return &pb.Result{Result: "You Lose\n"}, nil
	}

	//ini kertas
	if strings.ToLower(clientInput) == "kertas" && strings.ToLower(answer) == "gunting" {
		fmt.Printf("You Win\n")
		return &pb.Result{Result: "You Lose\n"}, nil
	} else if strings.ToLower(clientInput) == "kertas" && strings.ToLower(answer) == "kertas" {
		fmt.Printf("Draw\n")
		return &pb.Result{Result: "Draw\n"}, nil
	} else if strings.ToLower(clientInput) == "kertas" && strings.ToLower(answer) == "batu" {
		fmt.Printf("You Lose\n")
		return &pb.Result{Result: "You Win\n"}, nil
	} else {
		fmt.Printf("Invalid\n")
		return &pb.Result{Result: "Invalid Option"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

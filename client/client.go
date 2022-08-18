package main

import (
	"context"
	"log"

	pb "github.com/Kakashi944/Appointment_GRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAppointmentClient(conn)
	a := pb.Id{
		Id: 2,
	}
	bookList, err := client.GetAppointment(context.Background(), &a)
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("book list: %v", bookList)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"encoding/json"

	"github.com/Kakashi944/Appointment_GRPC/models"
	pb "github.com/Kakashi944/Appointment_GRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedAppointmentServer
}

func (s *server) NewAppointment(ctx context.Context, appoint *pb.Request) (*pb.Id, error) {
	log.Printf("request: %v", appoint)
	var err error

	var ident []string
	for i := range appoint.Identifier {
		identJson, _ := json.Marshal(models.Identifier{
			System: appoint.Identifier[i].System,
			Value:  appoint.Identifier[i].Value,
		})
		identString := string(identJson)
		ident = append(ident, identString)
	}

	text := models.Text{
		Status: appoint.Text.Status,
		Div:    appoint.Text.Div,
	}
	textJson, _ := json.Marshal(text)
	textString := string(textJson)

	appointment := models.Appointment{
		ResourceType: appoint.ResourceType,
		Text:         textString,
		Identifier:   ident,
	}
	id, err := models.InsertAppointment(appointment)
	if err != nil {
		panic(err)
	}

	return &pb.Id{Id: int32(id)}, err
}

func (s *server) GetAppointment(ctx context.Context, request *pb.Id) (*pb.Request, error) {

	fmt.Println("received : ", request.Id)
	ident := pb.Identifier{
		System: "http://example.org/sampleappointment-identifier",
		Value:  "456",
	}
	identifier := []*pb.Identifier{}
	identifier = append(identifier, &ident)

	return &pb.Request{
		ResourceType: "Appointment",
		Text: &pb.Text{
			Status: "generated",
			Div:    "Brian MRI results",
		},
		Identifier: identifier,
		Priority:   5,
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterAppointmentServer(s, &server{})
	log.Println("Server Started..........")
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}

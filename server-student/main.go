package main

import (
	"log"
	"net"

	"github.com/neecosanudo/platzi-grpc/database"
	"github.com/neecosanudo/platzi-grpc/server"
	"github.com/neecosanudo/platzi-grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	// Utilizamos gRPC por primera vez
	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	// reflection va a servir para poder proveer cierta meta-data a los clientes
	reflection.Register(s)

	// Levantamos el servidor
	err = s.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}

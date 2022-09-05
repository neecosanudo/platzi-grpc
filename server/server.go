package server

import (
	"context"

	"github.com/neecosanudo/platzi-grpc/models"
	"github.com/neecosanudo/platzi-grpc/repository"
	"github.com/neecosanudo/platzi-grpc/studentpb"
)

type Server struct {
	repo repository.Repository
	// Composition Over Inheritance: importante para que el servidor grpc admita este server
	studentpb.UnimplementedStudentServiceServer
}

// Constructor del servidor
func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

// Implementamos las funciones GetStudent y SetStudent
func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	// Traemos el estudiante que se encontró a través del repository
	student, err := s.repo.GetStudent(ctx, req.GetId())
	/**
		req.GetId() puede ser pasado como req.Id
		Ambos refieren al ID que declaramos en el archivo .proto:

		message GetStudentRequest {
	  	string id = 1;
		} */
	if err != nil {
		return nil, err
	}

	// &studentpb.Student viene del protobuffer que definimos.
	return &studentpb.Student{
		// student.Id, student.Name y student.Age, vienen del struct en models/
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	// Creamos el estudiante que vamos a enviar por el servidor
	student := &models.Student{ // struct que definimos en models/
		// req.GetId(), req.GetName() y req.GetAge(), vienen del protobuffer
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}

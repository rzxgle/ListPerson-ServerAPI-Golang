package main

import (
	"apidepartment/src/pb/department"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	department.DepartmentServiceServer
}

func (s *server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersoonServer) error {
	return nil
}

func main() {
	fmt.Println("starting grpc server")
	listener, err := net.Listen("tcp", ":9010")
	if err != nil {
		log.Fatalln("error on get listener. error: ", err)
	}

	s := grpc.NewServer()
	department.RegisterDepartmentServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on serve. error: ", err)
	}
}

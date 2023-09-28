package main

import (
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/config"
	"english/backend/internal/db"
	"english/backend/internal/services/svc-course/services"
	"english/backend/shared/errorlog"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.NewSVCConfig("course")
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		return
	}

	handler, err := db.Init(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		return
	}
	log.Println("Auth svc on ", cfg.Port)

	server := services.Server{
		H: handler,
	}

	grpcServer := grpc.NewServer()
	svccourse.RegisterCourseServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

package main

import (
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/config"
	"english/backend/internal/db"
	"english/backend/internal/services/svc-auth/services"
	"english/backend/shared/errorlog"
	sharedjwt "english/backend/shared/jwt"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.NewSVCConfig("auth")
	if err != nil {
		fmt.Println(err)
		return
	}

	handler, err := db.Init(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	jwt := sharedjwt.JwtWrapper{
		SecretKey:       cfg.JWTSecretKey,
		Issuer:          "auth-svc",
		ExpirationHours: 24,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		return
	}
	log.Println("Auth svc on ", cfg.Port)

	server := services.Server{
		H:   handler,
		JWT: jwt,
	}

	grpcServer := grpc.NewServer()
	svcauth.RegisterAuthServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

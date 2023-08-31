package auth

import (
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client svcauth.AuthServiceClient
}

func InitServiceClient(cfg *config.MainConfig) svcauth.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Print("Could not connect:\n", err)
	}

	return svcauth.NewAuthServiceClient(cc)
}

package client

import (
	svcclient "english/backend/api/proto/svc-client"
	"english/backend/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client svcclient.ClientServiceClient
}

func IniServiceClient(cfg *config.MainConfig) svcclient.ClientServiceClient {
	cc, err := grpc.Dial(cfg.ClientSvcURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:\n", err)
	}

	return svcclient.NewClientServiceClient(cc)
}

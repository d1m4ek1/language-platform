package teacher

import (
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/config"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client svccourse.CourseServiceClient
}

func InitServiceClient(cfg *config.MainConfig) svccourse.CourseServiceClient {
	cc, err := grpc.Dial(cfg.CourseSvcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:\n", err)
	}

	return svccourse.NewCourseServiceClient(cc)
}

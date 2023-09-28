package services

import (
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/internal/db"
)

type Server struct {
	H db.Handler
	svccourse.UnimplementedCourseServiceServer
}

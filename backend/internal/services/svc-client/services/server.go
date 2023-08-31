package services

import (
	svcclient "english/backend/api/proto/svc-client"
	"english/backend/internal/db"
)

type Server struct {
	H db.Handler
	svcclient.UnsafeClientServiceServer
}

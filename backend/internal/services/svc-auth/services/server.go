package services

import (
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/internal/db"
	sharedjwt "english/backend/shared/jwt"
)

type Server struct {
	H   db.Handler
	JWT sharedjwt.JwtWrapper
	svcauth.UnimplementedAuthServiceServer
}

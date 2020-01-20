package utils

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// DialGrpc dials a grpc service and panics if it could not connect.
func DialGrpc(service string) *grpc.ClientConn {
	grpcConn, err := grpc.Dial(service, grpc.WithInsecure())
	if err != nil {
		log.Panic().Err(err).
			Str("url", service).
			Msg("Cannot connect to GRPC service")
	}

	return grpcConn
}

package utils

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// DialGrpc dials a grpc service and panics if it could not connect.
func DialGrpc(service string, servicePort int32) *grpc.ClientConn {
	var address string = service
	if servicePort > -1 {
		address = fmt.Sprintf("%s:%d", address, servicePort)
	}

	grpcConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Panic().Err(err).
			Str("url", service).
			Msg("Cannot connect to GRPC service")
	}

	return grpcConn
}

package result

import (
	"bytes"
	"context"
	"io"

	"github.com/caspr-io/caspr/api/provisioning"
	provisioningapi "github.com/caspr-io/caspr/api/provisioning"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type Payload map[string]interface{}

func ReadPayload(reader io.Reader) Payload {
	payload := make(Payload)

	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(payload); err != nil {
		panic(err)
	}

	log.Logger.Info().Interface("payload", payload).Msg("Received payload to send")

	return payload
}

func (p Payload) Send(service string, subscription string) {
	grpcConn, err := grpc.Dial(service, grpc.WithInsecure())
	if err != nil {
		log.Panic().Err(err).
			Str("url", service).
			Msg("Cannot connect to provisioning-service")
	}

	provisioningServiceClient := provisioningapi.NewProvisioningServiceClient(grpcConn)

	byteBuffer := &bytes.Buffer{}

	encoder := yaml.NewEncoder(byteBuffer)
	defer encoder.Close()

	if err := encoder.Encode(p); err != nil {
		panic(err)
	}

	msg := &provisioning.ProvisioningResult{Subscription: subscription, Payload: byteBuffer.Bytes()}
	if _, err := provisioningServiceClient.Result(context.Background(), msg); err != nil {
		panic(err)
	}
}

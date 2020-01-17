package result

import (
	"context"
	"io"

	"github.com/caspr-io/caspr/api/provisioning"
	provisioningapi "github.com/caspr-io/caspr/api/provisioning"
	"github.com/caspr-io/caspr/internal/utils"
	"github.com/rs/zerolog/log"
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

func (p Payload) Send(service string, servicePort int32, subscription string) {
	grpcConn := utils.DialGrpc(service, servicePort)
	provisioningServiceClient := provisioningapi.NewProvisioningServiceClient(grpcConn)

	msg := &provisioning.ProvisioningResult{Subscription: subscription, Payload: utils.ToYaml(p)}
	if _, err := provisioningServiceClient.Result(context.Background(), msg); err != nil {
		panic(err)
	}
}

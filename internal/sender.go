package internal

import (
	"bufio"
	"bytes"
	"context"
	"os"

	"github.com/caspr-io/caspr-result/api/provisioning"
	"github.com/caspr-io/mu-kit/kit"
	"github.com/caspr-io/mu-kit/streaming"
	"gopkg.in/yaml.v3"
)

type Payload map[string]interface{}

func ReadPayload(reader *bufio.Reader) Payload {
	payload := make(Payload)

	decoder := yaml.NewDecoder(os.Stdin)
	if err := decoder.Decode(payload); err != nil {
		panic(err)
	}

	return payload
}

func (p Payload) Send() {
	stanConfig := &streaming.Config{}
	if err := kit.ReadConfig("", stanConfig); err != nil {
		panic(err)
	}

	river, err := streaming.NewRiver(stanConfig)
	if err != nil {
		panic(err)
	}

	byteBuffer := &bytes.Buffer{}

	encoder := yaml.NewEncoder(byteBuffer)
	defer encoder.Close()

	if err := encoder.Encode(p); err != nil {
		panic(err)
	}

	msg := &provisioning.ProvisioningResult{Payload: byteBuffer.Bytes()}
	if err := river.Publish(context.Background(), msg); err != nil {
		panic(err)
	}
}

module github.com/caspr-io/caspr

go 1.13

require (
	github.com/caspr-io/mu-kit v0.0.96
	github.com/caspr-io/yamlpath v0.0.1
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/spf13/cobra v1.0.0
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	gotest.tools/v3 v3.0.2
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
)

// replace github.com/caspr-io/mu-kit => ../mu-kit

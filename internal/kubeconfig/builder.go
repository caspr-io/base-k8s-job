package kubeconfig

import (
	"fmt"

	clusterapi "github.com/caspr-io/caspr/api/cluster"
	"github.com/caspr-io/caspr/internal/utils"
)

type Builder struct {
	configType      string
	credentialFile  string
	targetNamespace string
}

func NewBuilder(t string, file string, ns string) *Builder {
	return &Builder{
		configType:      t,
		credentialFile:  file,
		targetNamespace: ns,
	}
}

func (b *Builder) Build(output string) error {
	contents, err := utils.ReadYaml(b.credentialFile)
	if err != nil {
		return err
	}

	switch b.configType {
	case clusterapi.ClusterDetails_KUBERNETES.String():
		return utils.WriteYaml(output, b.kubernetesFormat(contents))
	case clusterapi.ClusterDetails_AWS.String():
		return utils.WriteYaml(output, b.awsFormat(contents))
	default:
		return fmt.Errorf("unsupported config type '%s'", b.configType)
	}
}

func (b *Builder) kubernetesFormat(contents map[string]interface{}) map[string]interface{} {
	return b.kubeConfig(contents, b.certificateKeyUser(contents))
}

func (b *Builder) awsFormat(contents map[string]interface{}) map[string]interface{} {
	return b.kubeConfig(contents, b.awsUser(contents))
}

func (b *Builder) certificateKeyUser(contents map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"client-certificate-data": contents["client-certificate"],
		"client-key-data":         contents["client-key"],
	}
}

func (b *Builder) awsUser(contents map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"exec": map[string]interface{}{
			"apiVersion": "client.authentication.k8s.io/v1alpha1",
			"args": []string{
				"token",
				"-i",
				contents["cluster-name"].(string),
			},
			"command": "aws-iam-authenticator",
			"env":     "null",
		},
	}
}

//nolint:gofmt
func (b *Builder) kubeConfig(contents map[string]interface{}, userData map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "v1",
		"clusters": []map[string]interface{}{
			map[string]interface{}{
				"cluster": map[string]interface{}{
					"certificate-authority-data": contents["certificate-authority"],
					"server":                     contents["server"],
				},
				"name": "kube-job-target",
			},
		},
		"contexts": []map[string]interface{}{
			map[string]interface{}{
				"context": map[string]interface{}{
					"cluster":   "kube-job-target",
					"namespace": b.targetNamespace,
					"user":      "target-credentials",
				},
				"name": "job-context",
			},
		},
		"current-context": "job-context",
		"kind":            "Config",
		"preferences":     map[string]interface{}{},
		"users": []map[string]interface{}{
			map[string]interface{}{
				"name": "target-credentials",
				"user": userData,
			},
		},
	}
}

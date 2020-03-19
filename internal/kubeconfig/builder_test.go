package kubeconfig

import (
	"testing"

	"gotest.tools/v3/assert"

	clusterapi "github.com/caspr-io/caspr/api/cluster"
)

func TestAwsUserWithoutRole(t *testing.T) {
	b := NewBuilder(clusterapi.ClusterGroup_AWS.String(), "/tmp/dummy", "dummy-ns")
	casprConfig := map[string]interface{}{
		"cluster": map[string]interface{}{
			"cluster-name": "TestAwsUserWithoutRole-cluster",
		},
	}

	awsConfig := b.awsUser(casprConfig)
	assert.Assert(t, awsConfig["exec"] != nil)
	awsConfigExec := awsConfig["exec"].(map[string]interface{})
	assert.Equal(t, awsConfigExec["apiVersion"], "client.authentication.k8s.io/v1alpha1")
	assert.Assert(t, awsConfigExec["args"] != nil)
	awsConfigExecArgs := awsConfigExec["args"].([]string)
	assert.Equal(t, len(awsConfigExecArgs), 3)
	assert.Equal(t, awsConfigExecArgs[2], "TestAwsUserWithoutRole-cluster")
}

func TestAwsUserWitRole(t *testing.T) {
	b := NewBuilder(clusterapi.ClusterGroup_AWS.String(), "/tmp/dummy", "dummy-ns")
	casprConfig := map[string]interface{}{
		"cluster": map[string]interface{}{
			"cluster-name": "TestAwsUserWithRole-cluster",
			"role-arn":     "TestAwsUserWithRole-role",
		},
	}

	awsConfig := b.awsUser(casprConfig)
	assert.Assert(t, awsConfig["exec"] != nil)
	awsConfigExec := awsConfig["exec"].(map[string]interface{})
	assert.Equal(t, awsConfigExec["apiVersion"], "client.authentication.k8s.io/v1alpha1")
	assert.Assert(t, awsConfigExec["args"] != nil)
	awsConfigExecArgs := awsConfigExec["args"].([]string)
	assert.Equal(t, len(awsConfigExecArgs), 5)
	assert.Equal(t, awsConfigExecArgs[2], "TestAwsUserWithRole-cluster")
	assert.Equal(t, awsConfigExecArgs[4], "TestAwsUserWithRole-role")
}

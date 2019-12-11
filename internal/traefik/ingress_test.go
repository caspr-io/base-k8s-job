package traefik

// import (
// 	"context"
// 	"testing"

// 	"github.com/caspr-io/mu-kit/kubernetes"
// 	"github.com/rs/zerolog/log"
// 	// "gitlab.com/caspr-io/application-service/db"
// 	"gotest.tools/v3/assert"
// )

// func TestCreateIngress(t *testing.T) {
// 	t.SkipNow() // TODO
// 	k8s, err := kubernetes.ConnectToKubernetes()
// 	assert.NilError(t, err)

// 	ctx := log.Logger.WithContext(context.Background())
// 	url, err := CreateIngressForServices(ctx, k8s, "app-appels-6", &db.Instance{
// 		Subscription: "foo",
// 		ApplicationObj: &db.Application{
// 			Vendor:     "acme",
// 			Name:       "medic",
// 			URLPattern: "http://<tenant>.127.0.0.1.xip.io/",
// 			Ingresses: []*db.Ingress{
// 				&db.Ingress{
// 					ServiceName: "nginx",
// 					Port:        80,
// 					TLS:         false,
// 				},
// 			},
// 		},
// 	}, map[string]string{})

// 	assert.NilError(t, err)

// 	assert.Equal(t, url, "http://foo.127.0.0.1.xip.io/")
// }

package traefik

// import (
// 	"context"
// 	"fmt"
// 	"net/url"
// 	"strings"

// 	"github.com/caspr-io/ingress/pkg/apis/traefik/v1alpha1"
// 	"github.com/caspr-io/ingress/pkg/client/clientset/versioned"
// 	"github.com/caspr-io/mu-kit/kubernetes"
// 	"github.com/pkg/errors"
// 	"github.com/rs/zerolog/log"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// )

// func CreateIngressForServices(ctx context.Context, k8s *kubernetes.K8s, namespace string, application *db.Instance, variables map[string]string) (*string, error) {
// 	log.Ctx(ctx).Info().Msg("Creating ingresses for provisioned application")

// 	replacedURL := createURL(application)

// 	logger := log.Ctx(ctx).With().Str("url", replacedURL).Logger()

// 	for _, svc := range application.ApplicationObj.Ingresses {
// 		logger.Info().Interface("service", svc).Msg("Creating ingress for service")

// 		if err := sanityCheck(k8s, namespace, svc); err != nil {
// 			return nil, err
// 		}

// 		ingressRoute := v1alpha1.NewIngressRoute(
// 			fmt.Sprintf("ingress-%s-%d", svc.ServiceName, svc.Port),
// 			DefineLabels(application, variables))

// 		if svc.TLS {
// 			ingressRoute.Spec.Entrypoints = []string{"websecure"}
// 			ingressRoute.Spec.TLS = &v1alpha1.IngressRouteTlsSpec{
// 				CertResolver: "default",
// 			}
// 		} else {
// 			ingressRoute.Spec.Entrypoints = []string{"web"}
// 			ingressRoute.Spec.TLS = nil
// 		}

// 		u, err := url.Parse(replacedURL)
// 		if err != nil {
// 			return nil, errors.Wrapf(err, "Could not parse URL for %s", application.URL)
// 		}

// 		host := u.Hostname()
// 		match := fmt.Sprintf("Host(`%s`)", host)

// 		if len(svc.PathPrefix) > 0 {
// 			match = fmt.Sprintf("%s && PathPrefix(`%s`)", match, svc.PathPrefix)
// 		}

// 		ingressRoute.Spec.Routes = []v1alpha1.IngressRouteRuleSpec{
// 			v1alpha1.IngressRouteRuleSpec{
// 				Kind:  "Rule",
// 				Match: match,
// 				Services: []v1alpha1.IngressRouteRuleServiceSpec{
// 					v1alpha1.IngressRouteRuleServiceSpec{
// 						Name: svc.ServiceName,
// 						Port: svc.Port,
// 					},
// 				},
// 			},
// 		}

// 		logger.Info().Interface("ingress", ingressRoute).Msg("Creating ingress")

// 		client, err := versioned.NewForConfig(k8s.Config)
// 		if err != nil {
// 			return nil, err
// 		}

// 		_, err = client.TraefikV1alpha1().IngressRoutes(namespace).Create(ingressRoute)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &replacedURL, nil
// }

// func sanityCheck(k8s *kubernetes.K8s, namespace string, ingress *db.Ingress) error {
// 	// Sanity check, does the service exist?
// 	service, err := k8s.CoreV1().Services(namespace).Get(ingress.ServiceName, metav1.GetOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	// Sanity check, does the service expose the port?
// 	portFound := false
// 	for _, p := range service.Spec.Ports {
// 		portFound = portFound || p.Port == ingress.Port
// 	}

// 	if !portFound {
// 		return fmt.Errorf("service %s in namespace %s does not expose port %d", ingress.ServiceName, namespace, ingress.Port)
// 	}

// 	return nil
// }

// func createURL(appInstance *db.Instance) string {
// 	urlTemplate := appInstance.ApplicationObj.URLPattern
// 	// TODO: Pass variables from tenant-service to application-service
// 	replacer := strings.NewReplacer(
// 		"<tenant>", appInstance.Subscription)
// 	// replacer := strings.NewReplacer(
// 	// 	"<vendor>", appInstance.ApplicationDefinition.VendorName,
// 	// 	"<application>", appInstance.ApplicationDefinition.ApplicationName,
// 	// 	"<tenant>", appInstance.TenantName)

// 	return replacer.Replace(urlTemplate)
// }

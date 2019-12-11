package traefik

import (
	"fmt"

	"github.com/caspr-io/caspr/pkg/apis/traefik/v1alpha1"
	"github.com/caspr-io/mu-kit/kubernetes"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Traefik struct {
	k8s *kubernetes.K8s
}

func NewTraefik() (*Traefik, error) {
	k8s, err := kubernetes.ConnectToKubernetes()
	if err != nil {
		return nil, err
	}

	return &Traefik{
		k8s: k8s,
	}, nil
}

func (t *Traefik) CreateIngress(ingress *Ingress) error {
	if err := checkNamespace(t.k8s, ingress.Namespace); err != nil {
		return err
	}

	if err := checkServiceAndPort(t.k8s, ingress); err != nil {
		return err
	}

	ingressRoute := v1alpha1.NewIngressRoute(
		fmt.Sprintf("ingress-%s-%d", ingress.Service, ingress.Port),
		ingress.Labels)

	if ingress.TLS {
		ingressRoute.Spec.Entrypoints = []string{"websecure"}
		ingressRoute.Spec.TLS = &v1alpha1.IngressRouteTLSSpec{
			CertResolver: "default",
		}
	} else {
		ingressRoute.Spec.Entrypoints = []string{"web"}
		ingressRoute.Spec.TLS = nil
	}

	// u, err := url.Parse(replacedURL)
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "Could not parse URL for %s", application.URL)
	// }

	// host := u.Hostname()
	// match := fmt.Sprintf("Host(`%s`)", host)

	// if len(svc.PathPrefix) > 0 {
	// 	match = fmt.Sprintf("%s && PathPrefix(`%s`)", match, svc.PathPrefix)
	// }
	return nil
}

func checkNamespace(k8s *kubernetes.K8s, namespace string) error {
	ns, err := k8s.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}

	log.Logger.Debug().Str("namespace", ns.Name).Msg("Namespace exists")

	return nil
}

func checkServiceAndPort(k8s *kubernetes.K8s, ingress *Ingress) error {
	svc, err := k8s.CoreV1().Services(ingress.Namespace).Get(ingress.Service, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// Sanity check, does the service expose the port?
	portFound := false
	for _, p := range svc.Spec.Ports {
		portFound = portFound || p.Port == ingress.Port
	}

	if !portFound {
		return fmt.Errorf("service %s in namespace %s does not expose port %d", ingress.Service, ingress.Namespace, ingress.Port)
	}

	return nil
}

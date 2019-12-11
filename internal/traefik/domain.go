package traefik

type Ingress struct {
	Namespace  string
	Service    string
	Port       int32
	TLS        bool
	Labels     map[string]string
	URL        string
	PathPrefix string
}

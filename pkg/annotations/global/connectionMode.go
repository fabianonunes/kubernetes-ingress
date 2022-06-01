package global

import (
	"fmt"
	"github.com/haproxytech/client-native/v3/models"

	"github.com/haproxytech/kubernetes-ingress/pkg/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/pkg/store"
)

type HTTPConnectionMode struct {
	name     string
	defaults *models.Defaults
}

func NewHTTPConnectionMode(n string, d *models.Defaults) *HTTPConnectionMode {
	return &HTTPConnectionMode{name: n, defaults: d}
}

func (a *HTTPConnectionMode) GetName() string {
	return a.name
}

func (a *HTTPConnectionMode) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)

	switch input {
	case
		"",
		"http-keep-alive",
		"http-server-close",
		"httpclose":
		break
	default:
		return fmt.Errorf("invalid http-connection-mode value '%s'", input)
	}
	a.defaults.HTTPConnectionMode = input
	return nil
}

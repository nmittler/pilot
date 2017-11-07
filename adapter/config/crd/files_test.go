package crd

import (
	"istio.io/pilot/model"
	"testing"
	"istio.io/pilot/adapter/config/memory"
)

var (
	allConfigs = []FileConfig{{
			meta: model.ConfigMeta{Type: model.DestinationPolicy.Type, Name: "circuit-breaker"},
			file: "testdata/cb-policy.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "timeout"},
			file: "testdata/timeout-route-rule.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "weighted"},
			file: "testdata/weighted-route.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "fault"},
			file: "testdata/fault-route.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "redirect"},
			file: "testdata/redirect-route.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "rewrite"},
			file: "testdata/rewrite-route.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "websocket"},
			file: "testdata/websocket-route.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.EgressRule.Type, Name: "google"},
			file: "testdata/egress-rule.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.DestinationPolicy.Type, Name: "egress-circuit-breaker"},
			file: "testdata/egress-rule-cb-policy.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.RouteRule.Type, Name: "egress-timeout"},
			file: "testdata/egress-rule-timeout-route-rule.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.IngressRule.Type, Name: "world"},
			file: "testdata/ingress-route-world.yaml.golden",
		}, {
			meta: model.ConfigMeta{Type: model.IngressRule.Type, Name: "foo"},
			file: "testdata/ingress-route-foo.yaml.golden",
		}}
)

func TestAllConfigs(t *testing.T) {
	mockStore := memory.Make(model.IstioConfigTypes)
	configStore := NewFileConfigStore(mockStore)

	for i := range allConfigs {
		input := allConfigs[i]
		configStore.CreateFromFile(input)
		_, exists := configStore.GetForFile(input)
		if (!exists) {
			t.Fatalf("missing config ", input)
		}
		// TODO(nmmittler): Compare meta? Do we care?
	}
}

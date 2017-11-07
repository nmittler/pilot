package discovery

import (
	"istio.io/pilot/proxy/envoy"
	"istio.io/pilot/adapter/config/memory"
	"istio.io/pilot/model"
	"istio.io/pilot/proxy"
	"istio.io/pilot/test/mock"
	proxyconfig "istio.io/api/proxy/v1/config"
	"time"
	"github.com/golang/protobuf/ptypes"
	"istio.io/pilot/adapter/serviceregistry/aggregate"
	"istio.io/pilot/platform"
)

var (
	defaultDiscoveryOptions = envoy.DiscoveryServiceOptions {
		Port : 8080,
		EnableProfiling: true,
		EnableCaching: true }
)

func makeMeshConfig() proxyconfig.MeshConfig {
	mesh := proxy.DefaultMeshConfig()
	mesh.MixerAddress = "istio-mixer.istio-system:9091"
	mesh.RdsRefreshDelay = ptypes.DurationProto(10 * time.Millisecond)
	return mesh
}

// mockController specifies a mock Controller for testing
type mockController struct{}

func (c *mockController) AppendServiceHandler(f func(*model.Service, model.Event)) error {
	return nil
}

func (c *mockController) AppendInstanceHandler(f func(*model.ServiceInstance, model.Event)) error {
	return nil
}

func (c *mockController) Run(<-chan struct{}) {}

func buildMockController() *aggregate.Controller {
	discovery1 := mock.NewDiscovery(
		map[string]*model.Service{
			mock.HelloService.Hostname:   mock.HelloService,
			mock.ExtHTTPService.Hostname: mock.ExtHTTPService,
		}, 2)

	discovery2 := mock.NewDiscovery(
		map[string]*model.Service{
			mock.WorldService.Hostname:    mock.WorldService,
			mock.ExtHTTPSService.Hostname: mock.ExtHTTPSService,
		}, 2)

	registry1 := aggregate.Registry{
		Name:             platform.ServiceRegistry("mockAdapter1"),
		ServiceDiscovery: discovery1,
		ServiceAccounts:  discovery1,
		Controller:       &mockController{},
	}

	registry2 := aggregate.Registry{
		Name:             platform.ServiceRegistry("mockAdapter2"),
		ServiceDiscovery: discovery2,
		ServiceAccounts:  discovery2,
		Controller:       &mockController{},
	}

	ctls := aggregate.NewController()
	ctls.AddRegistry(registry1)
	ctls.AddRegistry(registry2)

	return ctls
}

func NewDiscoveryService(configStore model.ConfigStore) (*envoy.DiscoveryService, error) {
	configController := memory.NewController(configStore)

	mesh := makeMeshConfig()

	// Use the mock discovery service with the hello and world services pre-loaded.
	mockDiscovery := mock.Discovery
	mockDiscovery.ClearErrors()

	environment := proxy.Environment{
		ServiceDiscovery: mockDiscovery,
		ServiceAccounts:  mockDiscovery,
		IstioConfigStore: model.MakeIstioStore(configStore),
		Mesh:             &mesh}

	serviceControllers := buildMockController()

	return envoy.NewDiscoveryService(
		serviceControllers,
		configController,
		environment,
		defaultDiscoveryOptions)
}
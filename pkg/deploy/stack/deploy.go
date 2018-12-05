package stack

import (
	"fmt"

	"github.com/rancher/rio/pkg/apply"
	"github.com/rancher/rio/pkg/istio/config"

	"github.com/rancher/rio/pkg/deploy/stack/input"
	"github.com/rancher/rio/pkg/deploy/stack/output"
	"github.com/rancher/rio/pkg/deploy/stack/populate"
	"github.com/rancher/rio/types/apis/rio.cattle.io/v1beta1"
)

func Deploy(
	namespace, space string,
	stack *v1beta1.Stack,
	configs []*v1beta1.Config,
	services []*v1beta1.Service,
	volumes []*v1beta1.Volume,
	routeSet []*v1beta1.RouteSet,
	externalServices []*v1beta1.ExternalService,
	injector *config.IstioInjector) error {

	input := &input.Stack{
		Namespace:        namespace,
		Space:            space,
		Stack:            stack,
		Services:         services,
		Volumes:          volumes,
		Configs:          configs,
		RouteSet:         routeSet,
		ExternalServices: externalServices,
	}

	output := output.NewDeployment()
	if !input.Stack.Spec.DisableMesh {
		output.Injectors = []apply.ConfigInjector{injector.Inject}
	}

	if err := populate.Populate(input, output); err != nil {
		return err
	}

	groupID := fmt.Sprintf("stack/%s/%s", namespace, stack.Name)
	return output.Deploy(namespace, groupID)
}

func Remove(namespace, space string, stack *v1beta1.Stack, injector *config.IstioInjector) error {
	return Deploy(namespace, space, stack, nil, nil, nil, nil, nil, injector)
}
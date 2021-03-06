// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type CustomResourceDefinitionWatcher interface {
	// watch cluster-scoped customresourcedefinition
	Watch(opts clients.WatchOpts) (<-chan CustomResourceDefinitionList, <-chan error, error)
}

type CustomResourceDefinitionClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(name string, opts clients.ReadOpts) (*CustomResourceDefinition, error)
	Write(resource *CustomResourceDefinition, opts clients.WriteOpts) (*CustomResourceDefinition, error)
	Delete(name string, opts clients.DeleteOpts) error
	List(opts clients.ListOpts) (CustomResourceDefinitionList, error)
	CustomResourceDefinitionWatcher
}

type customResourceDefinitionClient struct {
	rc clients.ResourceClient
}

func NewCustomResourceDefinitionClient(rcFactory factory.ResourceClientFactory) (CustomResourceDefinitionClient, error) {
	return NewCustomResourceDefinitionClientWithToken(rcFactory, "")
}

func NewCustomResourceDefinitionClientWithToken(rcFactory factory.ResourceClientFactory, token string) (CustomResourceDefinitionClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &CustomResourceDefinition{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base CustomResourceDefinition resource client")
	}
	return NewCustomResourceDefinitionClientWithBase(rc), nil
}

func NewCustomResourceDefinitionClientWithBase(rc clients.ResourceClient) CustomResourceDefinitionClient {
	return &customResourceDefinitionClient{
		rc: rc,
	}
}

func (client *customResourceDefinitionClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *customResourceDefinitionClient) Register() error {
	return client.rc.Register()
}

func (client *customResourceDefinitionClient) Read(name string, opts clients.ReadOpts) (*CustomResourceDefinition, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read("", name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*CustomResourceDefinition), nil
}

func (client *customResourceDefinitionClient) Write(customResourceDefinition *CustomResourceDefinition, opts clients.WriteOpts) (*CustomResourceDefinition, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(customResourceDefinition, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*CustomResourceDefinition), nil
}

func (client *customResourceDefinitionClient) Delete(name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete("", name, opts)
}

func (client *customResourceDefinitionClient) List(opts clients.ListOpts) (CustomResourceDefinitionList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List("", opts)
	if err != nil {
		return nil, err
	}
	return convertToCustomResourceDefinition(resourceList), nil
}

func (client *customResourceDefinitionClient) Watch(opts clients.WatchOpts) (<-chan CustomResourceDefinitionList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch("", opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	customresourcedefinitionChan := make(chan CustomResourceDefinitionList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				customresourcedefinitionChan <- convertToCustomResourceDefinition(resourceList)
			case <-opts.Ctx.Done():
				close(customresourcedefinitionChan)
				return
			}
		}
	}()
	return customresourcedefinitionChan, errs, nil
}

func convertToCustomResourceDefinition(resources resources.ResourceList) CustomResourceDefinitionList {
	var customResourceDefinitionList CustomResourceDefinitionList
	for _, resource := range resources {
		customResourceDefinitionList = append(customResourceDefinitionList, resource.(*CustomResourceDefinition))
	}
	return customResourceDefinitionList
}

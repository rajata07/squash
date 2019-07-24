// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type DebugAttachmentWatcher interface {
	// watch namespace-scoped Debugattachments
	Watch(namespace string, opts clients.WatchOpts) (<-chan DebugAttachmentList, <-chan error, error)
}

type DebugAttachmentClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*DebugAttachment, error)
	Write(resource *DebugAttachment, opts clients.WriteOpts) (*DebugAttachment, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (DebugAttachmentList, error)
	DebugAttachmentWatcher
}

type debugAttachmentClient struct {
	rc clients.ResourceClient
}

func NewDebugAttachmentClient(rcFactory factory.ResourceClientFactory) (DebugAttachmentClient, error) {
	return NewDebugAttachmentClientWithToken(rcFactory, "")
}

func NewDebugAttachmentClientWithToken(rcFactory factory.ResourceClientFactory, token string) (DebugAttachmentClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &DebugAttachment{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base DebugAttachment resource client")
	}
	return NewDebugAttachmentClientWithBase(rc), nil
}

func NewDebugAttachmentClientWithBase(rc clients.ResourceClient) DebugAttachmentClient {
	return &debugAttachmentClient{
		rc: rc,
	}
}

func (client *debugAttachmentClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *debugAttachmentClient) Register() error {
	return client.rc.Register()
}

func (client *debugAttachmentClient) Read(namespace, name string, opts clients.ReadOpts) (*DebugAttachment, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*DebugAttachment), nil
}

func (client *debugAttachmentClient) Write(debugAttachment *DebugAttachment, opts clients.WriteOpts) (*DebugAttachment, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(debugAttachment, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*DebugAttachment), nil
}

func (client *debugAttachmentClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *debugAttachmentClient) List(namespace string, opts clients.ListOpts) (DebugAttachmentList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToDebugAttachment(resourceList), nil
}

func (client *debugAttachmentClient) Watch(namespace string, opts clients.WatchOpts) (<-chan DebugAttachmentList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	debugattachmentsChan := make(chan DebugAttachmentList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				debugattachmentsChan <- convertToDebugAttachment(resourceList)
			case <-opts.Ctx.Done():
				close(debugattachmentsChan)
				return
			}
		}
	}()
	return debugattachmentsChan, errs, nil
}

func convertToDebugAttachment(resources resources.ResourceList) DebugAttachmentList {
	var debugAttachmentList DebugAttachmentList
	for _, resource := range resources {
		debugAttachmentList = append(debugAttachmentList, resource.(*DebugAttachment))
	}
	return debugAttachmentList
}

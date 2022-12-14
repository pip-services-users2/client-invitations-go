package build

import (
	clients1 "github.com/pip-services-users2/client-invitations-go/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type InvitationsClientFactory struct {
	*cbuild.Factory
}

func NewInvitationsClientFactory() *InvitationsClientFactory {
	c := &InvitationsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	cmdHttpClientDescriptor := cref.NewDescriptor("service-invitations", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-invitations", "client", "commandable-grpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-invitations", "client", "grpc", "*", "1.0")

	c.RegisterType(cmdHttpClientDescriptor, clients1.NewInvitationsHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewInvitationsCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewInvitationGrpcClientV1)

	return c
}

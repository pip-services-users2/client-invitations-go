package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-invitations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type invitationsGrpcClientV1Test struct {
	client  *version1.InvitationGrpcClientV1
	fixture *InvitationsClientFixtureV1
}

func newInvitationsGrpcClientV1Test() *invitationsGrpcClientV1Test {
	return &invitationsGrpcClientV1Test{}
}

func (c *invitationsGrpcClientV1Test) setup(t *testing.T) {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewInvitationGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewInvitationsClientFixtureV1(c.client)
}

func (c *invitationsGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newInvitationsGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}

func TestGrpcNotifyInvitation(t *testing.T) {
	c := newInvitationsGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestNotifyInvitation(t)
}

package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-invitations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type invitationsCommandableGrpcClientV1Test struct {
	client  *version1.InvitationsCommandableGrpcClientV1
	fixture *InvitationsClientFixtureV1
}

func newInvitationsCommandableGrpcClientV1Test() *invitationsCommandableGrpcClientV1Test {
	return &invitationsCommandableGrpcClientV1Test{}
}

func (c *invitationsCommandableGrpcClientV1Test) setup(t *testing.T) {
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

	c.client = version1.NewInvitationsCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewInvitationsClientFixtureV1(c.client)
}

func (c *invitationsCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newInvitationsCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}

func TestCommandableGrpcNotifyInvitation(t *testing.T) {
	c := newInvitationsCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestNotifyInvitation(t)
}

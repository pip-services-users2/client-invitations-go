package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-invitations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

var client *version1.InvitationGrpcClientV1
var fixture *InvitationsClientFixtureV1

func setup(t *testing.T) *InvitationsClientFixtureV1 {
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

	client = version1.NewInvitationGrpcClientV1()
	client.Configure(context.Background(), httpConfig)
	client.Open(context.Background(), "")

	fixture = NewInvitationsClientFixtureV1(client)

	return fixture
}

func teardown(t *testing.T) {
	client.Close(context.Background(), "")
}

func TestCrudOperations(t *testing.T) {
	fixture := setup(t)
	defer teardown(t)

	fixture.TestCrudOperations(t)
}

func TestNotifyInvitation(t *testing.T) {
	fixture := setup(t)
	defer teardown(t)

	fixture.TestNotifyInvitation(t)
}

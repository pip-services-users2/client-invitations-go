package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-invitations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type invitationsCommandableHttpClientV1Test struct {
	client  *version1.InvitationsCommandableHttpClientV1
	fixture *InvitationsClientFixtureV1
}

func newInvitationsCommandableHttpClientV1Test() *invitationsCommandableHttpClientV1Test {
	return &invitationsCommandableHttpClientV1Test{}
}

func (c *invitationsCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewInvitationsHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewInvitationsClientFixtureV1(c.client)
}

func (c *invitationsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newInvitationsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}

func TestCommandableHttpNotifyInvitation(t *testing.T) {
	c := newInvitationsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestNotifyInvitation(t)
}

package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type InvitationsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewInvitationsHttpClientV1() *InvitationsCommandableHttpClientV1 {
	return NewInvitationsHttpClientV1WithConfig(nil)
}

func NewInvitationsHttpClientV1WithConfig(config *cconf.ConfigParams) *InvitationsCommandableHttpClientV1 {
	c := &InvitationsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/invitations"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *InvitationsCommandableHttpClientV1) GetInvitations(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*InvitationV1], err error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_invitations", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*InvitationV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*InvitationV1]](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) GetInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "get_invitation_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) CreateInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation", invitation,
	)

	res, err := c.CallCommand(ctx, "create_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) DeleteInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "delete_invitation_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) ActivateInvitations(ctx context.Context, correlationId string, email string, userId string) (result []*InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"email", email,
		"user_id", userId,
	)

	res, err := c.CallCommand(ctx, "activate_invitations", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) ApproveInvitation(ctx context.Context, correlationId string, invitationId string, role string) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
		"role", role,
	)

	res, err := c.CallCommand(ctx, "deny_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) DenyInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "resend_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) ResendInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "resend_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableHttpClientV1) NotifyInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) error {
	params := data.NewAnyValueMapFromTuples(
		"invitation", invitation,
	)

	_, err := c.CallCommand(ctx, "notify_invitation", correlationId, params)

	return err
}

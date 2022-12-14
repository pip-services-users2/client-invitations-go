package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type InvitationsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewInvitationsCommandableGrpcClientV1() *InvitationsCommandableGrpcClientV1 {
	return NewInvitationsCommandableGrpcClientV1WithConfig(nil)
}

func NewInvitationsCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *InvitationsCommandableGrpcClientV1 {
	c := &InvitationsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/invitations"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *InvitationsCommandableGrpcClientV1) GetInvitations(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*InvitationV1], err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.get_invitations")
	defer timing.EndTiming(ctx, err)

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

func (c *InvitationsCommandableGrpcClientV1) GetInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.get_invitation_by_id")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "get_invitation_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableGrpcClientV1) CreateInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.create_invitation")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation", invitation,
	)

	res, err := c.CallCommand(ctx, "create_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableGrpcClientV1) DeleteInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.delete_invitation_by_id")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "delete_invitation_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableGrpcClientV1) ActivateInvitations(ctx context.Context, correlationId string, email string, userId string) (result []*InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.activate_invitations")
	defer timing.EndTiming(ctx, err)

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

func (c *InvitationsCommandableGrpcClientV1) ApproveInvitation(ctx context.Context, correlationId string, invitationId string, role string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.deny_invitation")
	defer timing.EndTiming(ctx, err)

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

func (c *InvitationsCommandableGrpcClientV1) DenyInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.resend_invitation")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "resend_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableGrpcClientV1) ResendInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.resend_invitation")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation_id", invitationId,
	)

	res, err := c.CallCommand(ctx, "resend_invitation", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*InvitationV1](res, correlationId)
}

func (c *InvitationsCommandableGrpcClientV1) NotifyInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) (err error) {
	timing := c.Instrument(ctx, correlationId, "invitations.notify_invitation")
	defer timing.EndTiming(ctx, err)

	params := data.NewAnyValueMapFromTuples(
		"invitation", invitation,
	)

	_, err = c.CallCommand(ctx, "notify_invitation", correlationId, params)

	return err
}

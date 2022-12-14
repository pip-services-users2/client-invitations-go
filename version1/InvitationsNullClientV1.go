package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type InvitationsNullClientV1 struct {
}

func (c *InvitationsNullClientV1) GetInvitations(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*InvitationV1], err error) {
	return *data.NewEmptyDataPage[*InvitationV1](), nil
}

func (c *InvitationsNullClientV1) GetInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) CreateInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) (result *InvitationV1, err error) {
	return invitation, nil
}

func (c *InvitationsNullClientV1) DeleteInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) ActivateInvitations(ctx context.Context, correlationId string, email string, userId string) (result []*InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) ApproveInvitation(ctx context.Context, correlationId string, invitationId string, role string) (result *InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) DenyInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) ResendInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error) {
	return nil, nil
}

func (c *InvitationsNullClientV1) NotifyInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) error {
	return nil
}

package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IInvitationsClientV1 interface {
	GetInvitations(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*InvitationV1], err error)

	GetInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error)

	CreateInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) (result *InvitationV1, err error)

	DeleteInvitationById(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error)

	ActivateInvitations(ctx context.Context, correlationId string, email string, userId string) (result []*InvitationV1, err error)

	ApproveInvitation(ctx context.Context, correlationId string, invitationId string, role string) (result *InvitationV1, err error)

	DenyInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error)

	ResendInvitation(ctx context.Context, correlationId string, invitationId string) (result *InvitationV1, err error)

	NotifyInvitation(ctx context.Context, correlationId string, invitation *InvitationV1) error
}

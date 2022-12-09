package version1

import (
	"context"

	"github.com/pip-services-users2/client-invitations-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type InvitationGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewInvitationGrpcClientV1() *InvitationGrpcClientV1 {
	return &InvitationGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("invitations_v1.Invitations"),
	}
}

func (c *InvitationGrpcClientV1) GetInvitations(ctx context.Context, correlationId string,
	filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*InvitationV1], err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.get_invitations")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.InvitationPageReply)
	err = c.CallWithContext(ctx, "get_invitations", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*InvitationV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*InvitationV1](), err
	}

	result = toInvitationPage(reply.Page)

	return result, nil
}

func (c *InvitationGrpcClientV1) GetInvitationById(ctx context.Context, correlationId string,
	invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.get_invitation_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationIdRequest{
		CorrelationId: correlationId,
		InvitationId:  invitationId,
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "get_invitation_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) CreateInvitation(ctx context.Context, correlationId string,
	invitation *InvitationV1) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.create_invitation")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationObjectRequest{
		CorrelationId: correlationId,
		Invitation:    fromInvitation(invitation),
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "create_invitation", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) DeleteInvitationById(ctx context.Context, correlationId string,
	invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.delete_invitation_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationIdRequest{
		CorrelationId: correlationId,
		InvitationId:  invitationId,
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "delete_invitation_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) ActivateInvitations(ctx context.Context, correlationId string,
	email string, userId string) (result []*InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.activate_invitations")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationActivateRequest{
		CorrelationId: correlationId,
		Email:         email,
		UserId:        userId,
	}

	reply := new(protos.InvitationListReply)
	err = c.CallWithContext(ctx, "activate_invitations", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitationList(reply.Invitations)

	return result, nil
}

func (c *InvitationGrpcClientV1) ApproveInvitation(ctx context.Context, correlationId string,
	invitationId string, role string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.approve_invitation")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationApproveRequest{
		CorrelationId: correlationId,
		InvitationId:  invitationId,
		Role:          role,
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "approve_invitation", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) DenyInvitation(ctx context.Context, correlationId string,
	invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.deny_invitation")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationIdRequest{
		CorrelationId: correlationId,
		InvitationId:  invitationId,
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "deny_invitation", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) ResendInvitation(ctx context.Context, correlationId string,
	invitationId string) (result *InvitationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.resend_invitation")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationIdRequest{
		CorrelationId: correlationId,
		InvitationId:  invitationId,
	}

	reply := new(protos.InvitationObjectReply)
	err = c.CallWithContext(ctx, "resend_invitation", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toInvitation(reply.Invitation)

	return result, nil
}

func (c *InvitationGrpcClientV1) NotifyInvitation(ctx context.Context, correlationId string,
	invitation *InvitationV1) (err error) {
	timing := c.Instrument(ctx, correlationId, "invitations_v1.notify_invitation")
	defer timing.EndTiming(ctx, err)

	req := &protos.InvitationObjectRequest{
		CorrelationId: correlationId,
		Invitation:    fromInvitation(invitation),
	}

	reply := new(protos.InvitationEmptyReply)
	err = c.CallWithContext(ctx, "notify_invitation", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

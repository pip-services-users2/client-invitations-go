package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-invitations-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromInvitation(invitation *InvitationV1) *protos.Invitation {
	if invitation == nil {
		return nil
	}

	obj := &protos.Invitation{
		Id:           invitation.Id,
		Action:       invitation.Action,
		OrgId:        invitation.OrgId,
		OrgName:      invitation.OrgName,
		Role:         invitation.Role,
		Language:     invitation.Language,
		CreateTime:   convert.StringConverter.ToString(invitation.CreateTime),
		CreatorId:    invitation.CreatorId,
		CreatorName:  invitation.CreatorName,
		InviteeId:    invitation.InviteeId,
		InviteeName:  invitation.InviteeName,
		InviteeEmail: invitation.InviteeEmail,
		InviteeAbout: invitation.InviteeAbout,
		SentTime:     convert.StringConverter.ToString(invitation.SentTime),
		ExpireTime:   convert.StringConverter.ToString(invitation.ExpireTime),
	}

	return obj
}

func toInvitation(obj *protos.Invitation) *InvitationV1 {
	if obj == nil {
		return nil
	}

	invitation := &InvitationV1{
		Id:           obj.Id,
		Action:       obj.Action,
		OrgId:        obj.OrgId,
		OrgName:      obj.OrgName,
		Role:         obj.Role,
		Language:     obj.Language,
		CreateTime:   convert.DateTimeConverter.ToDateTime(obj.CreateTime),
		CreatorId:    obj.CreatorId,
		CreatorName:  obj.CreatorName,
		InviteeId:    obj.InviteeId,
		InviteeName:  obj.InviteeName,
		InviteeEmail: obj.InviteeEmail,
		InviteeAbout: obj.InviteeAbout,
		SentTime:     convert.DateTimeConverter.ToDateTime(obj.SentTime),
		ExpireTime:   convert.DateTimeConverter.ToDateTime(obj.ExpireTime),
	}

	return invitation
}

func fromInvitationPage(page data.DataPage[*InvitationV1]) *protos.InvitationPage {
	obj := &protos.InvitationPage{
		Total: int64(page.Total),
		Data:  make([]*protos.Invitation, len(page.Data)),
	}

	for i, invitation := range page.Data {
		obj.Data[i] = fromInvitation(invitation)
	}

	return obj
}

func toInvitationPage(obj *protos.InvitationPage) data.DataPage[*InvitationV1] {
	invitations := make([]*InvitationV1, len(obj.Data))

	for i, v := range obj.Data {
		invitations[i] = toInvitation(v)
	}

	page := data.NewDataPage(invitations, int(obj.Total))

	return *page
}

func fromInvitationList(invitations []*InvitationV1) []*protos.Invitation {
	if invitations == nil {
		return nil
	}

	data := make([]*protos.Invitation, len(invitations))

	for i, v := range invitations {
		data[i] = fromInvitation(v)
	}

	return data
}

func toInvitationList(obj []*protos.Invitation) []*InvitationV1 {
	if obj == nil {
		return nil
	}

	invitations := make([]*InvitationV1, len(obj))

	for i, v := range obj {
		invitations[i] = toInvitation(v)
	}

	return invitations
}

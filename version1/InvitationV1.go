package version1

import (
	"time"
)

type InvitationV1 struct {
	/* Identification */
	Id string `json:"id"`

	/* Invitation info */
	Action   string `json:"action"`
	OrgId    string `json:"org_id"`
	OrgName  string `json:"org_name"`
	Role     string `json:"role"`
	Language string `json:"language"`

	/* Creator info */
	CreateTime  time.Time `json:"create_time"`
	CreatorName string    `json:"creator_name"`
	CreatorId   string    `json:"creator_id"`

	/* Invitee info */
	InviteeName  string `json:"invitee_name"`
	InviteeEmail string `json:"invitee_email"`
	InviteeId    string `json:"invitee_id"`
	InviteeAbout string `json:"invitee_about"`

	/* Timing */
	SentTime   time.Time `json:"sent_time"`
	ExpireTime time.Time `json:"expire_time"`
}

func EmptyInvitationV1() *InvitationV1 {
	return &InvitationV1{}
}

func NewInvitationV1(id string, action string, orgId string, orgName string,
	role string, inviteeId string, inviteeName string, inviteeEmail string) *InvitationV1 {
	return &InvitationV1{
		Id:           id,
		Action:       action,
		OrgId:        orgId,
		OrgName:      orgName,
		CreateTime:   time.Now(),
		InviteeId:    inviteeId,
		InviteeName:  inviteeName,
		InviteeEmail: inviteeEmail,
	}
}

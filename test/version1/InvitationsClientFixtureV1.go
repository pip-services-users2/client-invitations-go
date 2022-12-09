package test_version1

import (
	"context"
	"testing"
	"time"

	"github.com/pip-services-users2/client-invitations-go/version1"
	"github.com/stretchr/testify/assert"
)

type InvitationsClientFixtureV1 struct {
	Client version1.IInvitationsClientV1
}

var INVITATION1 = &version1.InvitationV1{
	Id:           "1",
	Action:       "activate",
	OrgId:        "1",
	Role:         "manager",
	CreateTime:   time.Now(),
	CreatorId:    "1",
	InviteeEmail: "test@somewhere.com",
}
var INVITATION2 = &version1.InvitationV1{
	Id:           "2",
	Action:       "activate",
	OrgId:        "1",
	CreateTime:   time.Now(),
	CreatorId:    "1",
	InviteeEmail: "test2@somewhere.com",
}
var INVITATION3 = &version1.InvitationV1{
	Id:           "4",
	Action:       "approve",
	OrgId:        "1",
	CreateTime:   time.Now(),
	CreatorId:    "1",
	InviteeEmail: "test2@somewhere.com",
}

func NewInvitationsClientFixtureV1(client version1.IInvitationsClientV1) *InvitationsClientFixtureV1 {
	return &InvitationsClientFixtureV1{
		Client: client,
	}
}

func (c *InvitationsClientFixtureV1) clear() {
	page, _ := c.Client.GetInvitations(context.Background(), "", nil, nil)

	for _, invitation := range page.Data {
		c.Client.DeleteInvitationById(context.Background(), "", invitation.Id)
	}
}

func (c *InvitationsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one invitation
	invitation, err := c.Client.CreateInvitation(context.Background(), "", INVITATION1)
	assert.Nil(t, err)

	assert.NotNil(t, invitation)
	assert.Equal(t, invitation.CreatorId, INVITATION1.CreatorId)
	assert.Equal(t, invitation.OrgId, INVITATION1.OrgId)
	assert.Equal(t, invitation.InviteeEmail, INVITATION1.InviteeEmail)

	invitation1 := invitation

	// Create another invitation
	invitation, err = c.Client.CreateInvitation(context.Background(), "", INVITATION2)
	assert.Nil(t, err)

	assert.NotNil(t, invitation)
	assert.Equal(t, invitation.CreatorId, INVITATION2.CreatorId)
	assert.Equal(t, invitation.OrgId, INVITATION2.OrgId)
	assert.Equal(t, invitation.InviteeEmail, INVITATION2.InviteeEmail)

	//invitation2 := invitation

	// Get all invitations
	page, err1 := c.Client.GetInvitations(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Delete invitation
	invitation, err = c.Client.DeleteInvitationById(context.Background(), "", invitation1.Id)
	assert.Nil(t, err)

	// Try to get deleted invitation
	invitation, err = c.Client.GetInvitationById(context.Background(), "", invitation1.Id)
	assert.Nil(t, err)

	assert.Nil(t, invitation)
}

func (c *InvitationsClientFixtureV1) TestNotifyInvitation(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one invitation
	err := c.Client.NotifyInvitation(context.Background(), "", INVITATION3)
	assert.Nil(t, err)
}

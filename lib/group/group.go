package group

import (
	"github.com/gofrs/uuid"

	"github.com/julianknodt/group/lib/integration"
	"github.com/julianknodt/group/lib/user"
)

type Uuid string

const InboxSize = 128

type Group struct {
	// Unique id for this group
	Uid uuid.UUID

	// Group that this is a subgroup of
	Parent *Group

	// Name of this group
	Name string

	// Named bundle of groups
	Subgroups map[string][]*Group

	// Permissions for a specific named bundle
	SubgroupPermissions map[string]*GroupPermission

	// Named bundle of users
	Users map[string][]*user.User

	// Permissions for a specific named user
	UserPermissions map[string]*user.UserPermission

	// Settings for this group
	// Eventually migrate this to an int field or bit array or something

	IsPubliclyVisible bool
	CanDeleteMessages bool

	// Will think of more...

	Inbox chan<- interface{}
	inbox <-chan interface{}

	Integrations           map[string][]*integration.Integration
	IntegrationPermissions map[string]*integration.IntegrationPermission
}

func NewRootGroup() *Group {
	inbox := make(chan interface{}, InboxSize)
	return &Group{
		Uid:    uuid.Must(uuid.NewV4()),
		Parent: nil,
		Name:   "World",

		Subgroups:           make(map[string][]*Group),
		SubgroupPermissions: make(map[string]*GroupPermission),

		Users:           make(map[string][]*user.User),
		UserPermissions: make(map[string]*user.UserPermission),

		IsPubliclyVisible: true,
		CanDeleteMessages: false,

		Inbox: inbox,
		inbox: inbox,
	}
}

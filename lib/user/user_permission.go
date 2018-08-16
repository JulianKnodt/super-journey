package user

// Represents permissions for a Group
type UserPermission struct {
	// Eventually move these to an int or bit field

	CanRead               bool
	CanWrite              bool
	CanEdit               bool
	CanDelete             bool
	CanAddIntegrations    bool
	CanDeleteIntegrations bool

	// maybe think of more
}

package gh

type AccountType string

const (
	AccountTypeOrg  AccountType = "Organization"
	AccountTypeUser AccountType = "User"
)

// Account is a GitHub account.
type Account struct {
	// Account id
	ID int
	// Account handle (some_user).
	Login string
	// Account name (Some User).
	Name string
	// Account type (Organization or User).
	Type AccountType
}

// User is a GitHub user.
type User struct {
	// User id
	ID int
	// User login (some_user).
	Login string
	// User name (Some User).
	Name string
	// The orgs the user is a member of.
	Orgs []*Account
	// The users' preferred git protocol.
	GitProtocol Protocol
}

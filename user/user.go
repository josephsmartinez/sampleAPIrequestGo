package user

// User data for url creatation
type User struct {
	FirstName    string
	LastName     string
	Password     string
	APIKey       string
	StreetNumber string
	StreetName   string
	City         string
	State        string
}

// NewUser -
func NewUser() *User {
	return &User{"", "", "", "", "", "", "", ""}
}

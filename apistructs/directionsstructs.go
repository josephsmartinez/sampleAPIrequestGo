package apistructs

// User data for url creatation
type User struct {
	FirstName string
	LastName  string
	Password  string
	Address   string
	Address2  string
	City      string
	State     string
}

// BlankNewUser -
func BlankNewUser() *User {
	return &User{"", "", "", "", "", "", ""}
}

// CustomNewUser - Example constructor good for handle unknown struct data initialization
func CustomNewUser(fname string, lname string, pass string, add string, add2 string, city string, state string) *User {
	user := User{
		FirstName: fname,
		LastName:  lname,
		Password:  pass,
		Address:   add,
		Address2:  add2,
		City:      city,
		State:     state}
	return &user
}

// AddressFormatter -
func (user *User) AddressFormatter() {

}

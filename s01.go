package structs

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

type User struct {
	firstName string
	lastName  string
}

func New() UserInterface {
	return &User{}
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func ResetUser(user User) {
	user.SetFirstName("")
	user.SetLastName("")
}

func IsUser(input interface{}) bool {
	return input.(UserInterface) != nil
}

func ProcessUser(input User) string {
	input.SetFirstName("John")
	return input.FullName()
}

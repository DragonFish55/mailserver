package main



type UserLogin struct {
	Username string
	Password string
}

type UserReg struct {
	Username string
	Password string
	Confpass string
	Fname    string
	Lname    string
	Bday     string
}

/*
type Email struct {
	sender   string
	receiver string
	content  string
	date     string
	favorite bool
	mailtype []MailType
}
*/

type MailType interface {

	getCapacity() int
	setCapacity()

}

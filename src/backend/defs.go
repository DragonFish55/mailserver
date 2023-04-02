package main



type UserLogin struct {
	username string
	password string
}

type UserReg struct {
	username string
	password string
	confpass string
	fname    string
	lname    string
	bday     string
}

type Email struct {
	sender   string
	receiver string
	content  string
	datetime string
	favorite bool
	mailtype []MailType
}

type MailBox struct {
	username string
	Email []emails


}

type MailType interface {

	getCapacity() int
	setCapacity()

}

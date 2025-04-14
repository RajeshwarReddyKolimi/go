package chatModels

type User struct {
	Id     int
	Name   string
	Gender string
}

type Message struct {
	Id       int
	Text     string
	Sender   User
	Receiver User
	Time     string
}

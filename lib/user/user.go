package user

type Uuid string

type User struct {
	Uid Uuid

	Inbox <-chan interface{}
}

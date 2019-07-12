package interfaces

type IUserListInputPort interface {
	Handle(string)
}

type IUserListOutputPort interface {
	Emit([]string)
}

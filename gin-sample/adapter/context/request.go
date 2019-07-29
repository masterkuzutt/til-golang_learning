package context

type UserListRequest struct {
	Username string
}

type RegisterRequest struct {
	Username string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}

package service

type (
	UserResponse struct {
		UserId   int    `json:"userid"`
		FullName string `json:"fullname"`
	}

	UsersResponse []UserResponse

	UserService interface {
		GetUser(uid int) (*UserResponse, error)
	}
)

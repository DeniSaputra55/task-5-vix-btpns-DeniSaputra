package user

import "github.com/DeniSaputra55/task-5-vix-btpns-DeniSaputra/models"

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponseWithToken struct {
	UserResponse
	Token string `json:"token"`
}

func FormatUserResponse(user models.User, token string) interface{} {
	var formatter interface{}

	if token == "" {
		formatter = UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	} else {
		userResponse := UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		formatter = UserResponseWithToken{
			UserResponse: userResponse,
			Token:        token,
		}
	}

	return formatter
}
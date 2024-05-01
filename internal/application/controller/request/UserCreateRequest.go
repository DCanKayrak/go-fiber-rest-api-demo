package request

import "golang-rest-api-demo/internal/application/handler/user"

type UserCreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
}

func (req *UserCreateRequest) ToCommand() user.Command {
	return user.Command{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Age:       req.Age,
	}
}

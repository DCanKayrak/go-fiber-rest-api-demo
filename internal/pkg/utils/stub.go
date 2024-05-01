package utils

import "golang-rest-api-demo/internal/domain"

func GetUserStub() []*domain.User {
	return []*domain.User{
		{
			Id:        "1",
			FirstName: "Danyal Can",
			LastName:  "KAYRAK",
			Email:     "dancankan@gmail.com",
			Password:  "1234",
			Age:       20,
		},
		{
			Id:        "2",
			FirstName: "Danyal Can",
			LastName:  "KAYRAK",
			Email:     "dancankan@gmail.com",
			Password:  "1234",
			Age:       20,
		},
	}
}

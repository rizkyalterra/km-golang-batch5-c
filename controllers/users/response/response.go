package response

import "clean_architecture/entities"

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromEntities(data *entities.User) Response {
	return Response{
		ID:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		Token: data.Token,
	}
}


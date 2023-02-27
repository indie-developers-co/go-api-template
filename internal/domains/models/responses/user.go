package responses

import "gitlab.com/indie-developers/go-api-echo-template/internal/domains/models"

type GetUsersResponse struct {
	models.User
}

func ConvertToGetUsersResponseStruct(users []models.User) []GetUsersResponse {
	response := make([]GetUsersResponse, len(users))
	for _, user := range users {
		response = append(response, GetUsersResponse{
			User: user,
		})
	}

	return response
}

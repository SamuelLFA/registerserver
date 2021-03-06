package dto

import "github.com/samuellfa/registerserver/src/models"

type CreatePersonRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=255"`
	Document string `json:"document" binding:"required,min=1,max=255"`
	Email    string `json:"email" binding:"required,email,min=1,max=255"`
}

type createPersonResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

type findPersonResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (request *CreatePersonRequest) ToModel() *models.Person {
	return &models.Person{
		Name:     request.Name,
		Document: request.Document,
		Email:    request.Email,
	}
}

func NewCreatePersonResponse(person models.Person) createPersonResponse {
	return createPersonResponse{
		Id:       person.Id,
		Name:     person.Name,
		Document: person.Document,
		Email:    person.Email,
	}
}

func NewFindAllPersonResponse(people []models.Person) []findPersonResponse {
	var peopleDTO []findPersonResponse
	for _, person := range people {
		personDTO := findPersonResponse{
			Id:       person.Id,
			Name:     person.Name,
			Document: person.Document,
			Email:    person.Email,
		}
		peopleDTO = append(peopleDTO, personDTO)
	}
	return peopleDTO
}

func NewFindPersonResponse(person models.Person) findPersonResponse {
	personDTO := findPersonResponse{
		Id:       person.Id,
		Name:     person.Name,
		Document: person.Document,
		Email:    person.Email,
	}
	return personDTO
}

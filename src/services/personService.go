package services

import (
	"github.com/gin-gonic/gin"
	"github.com/samuellfa/registerserver/src/controllers/dto"
	"github.com/samuellfa/registerserver/src/errors"
	"github.com/samuellfa/registerserver/src/models"
	"gorm.io/gorm"
)

type PersonService struct {
	db *gorm.DB
}

func NewPersonService(db *gorm.DB) *PersonService {
	return &PersonService{db}
}

func (service *PersonService) Create(request dto.CreatePersonRequest, ctx *gin.Context) (*models.Person, errors.APIError) {
	person := request.ToModel()
	var validationError errors.ValidationError

	var personWithSameEmail models.Person
	if service.db.First(&personWithSameEmail, models.Person{Email: request.Email}); personWithSameEmail.Id != "" {
		validationError.AddError("email already used")
	}

	var personWithSameDocument models.Person
	if service.db.First(&personWithSameDocument, models.Person{Document: request.Document}); personWithSameDocument.Id != "" {
		validationError.AddError("document already used")
	}

	if !validationError.IsEmpty() {
		return nil, &validationError
	}

	if e := service.db.Save(&person).Error; e != nil {
		return nil, errors.NewInternalServerError("something was wrong")
	}

	return person, nil
}

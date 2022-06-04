package services

import (
	"github.com/gin-gonic/gin"
	dto "github.com/samuellfa/registerserver/src/controllers/dto"
	e "github.com/samuellfa/registerserver/src/errors"
	m "github.com/samuellfa/registerserver/src/models"
	"gorm.io/gorm"
)

type PersonService struct {
	db *gorm.DB
}

func NewPersonService(db *gorm.DB) *PersonService {
	return &PersonService{db}
}

func (service *PersonService) Create(request dto.CreatePersonRequest, ctx *gin.Context) (*m.Person, e.APIError) {
	person := request.ToModel()
	var validationError e.ValidationError

	var personWithSameEmail m.Person
	if service.db.First(&personWithSameEmail, m.Person{Email: request.Email}); personWithSameEmail.Id != "" {
		validationError.AddError(e.ValidationErrorField{Field: "email", Message: "value already used"})
	}

	var personWithSameDocument m.Person
	if service.db.First(&personWithSameDocument, m.Person{Document: request.Document}); personWithSameDocument.Id != "" {
		validationError.AddError(e.ValidationErrorField{Field: "document", Message: "value already used"})
	}

	if !validationError.IsEmpty() {
		return nil, &validationError
	}

	if err := service.db.Save(&person).Error; err != nil {
		return nil, e.NewInternalServerError("something was wrong")
	}

	return person, nil
}

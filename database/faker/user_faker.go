package faker

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/oktadafa/gotoko/app/models"
	"gorm.io/gorm"
)


func UserFaker(db *gorm.DB) *models.User{
	return &models.User{
		ID: uuid.New().String(),
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Passoword: "Okta54321",
		RememberToken: "",
		CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	}
}
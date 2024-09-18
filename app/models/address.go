package models

import "time"

type Address struct{
		ID string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name string `gorm:"size:100;not null"`
	User User
	UserID string `gorm:"size:36;not null;index"`
	IsPrimary bool
	CityId string `gorm:"size:100;"`
	ProvinciId string `gorm:"size:100;"`
	Address1 string `gorm:"size:255;"`
	Address2 string `gorm:"size:255;"`
	Phone string `gorm:"size:100;"`
	Email string `gorm:"size:100;"`
	PostCode string `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
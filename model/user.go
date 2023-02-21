package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	// JSON bindin for the Password field is  "-". This ensures that the user's password is not returned in the JSON response.
	Password string `gorm:"size:255;not null;" json:"-"`
	Entries  []Entry
}

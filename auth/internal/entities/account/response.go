package account

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	AccountID          uuid.UUID `json:"accountID" gorm:"primary_key"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	IsConfirmed        bool      `json:"isConfirmed"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

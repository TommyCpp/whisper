package model

import (
	"github.com/satori/go.uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

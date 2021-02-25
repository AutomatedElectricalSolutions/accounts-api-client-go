package models

import "github.com/google/uuid"

type User struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func (u *User) GetUUID() uuid.UUID {
	id, _ := uuid.Parse(u.ID)
	return id
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// bson, json and validate tags should not have spaces after colons
	ID           primitive.ObjectID `bson:"_id,omitempty"` // "omitempty" allows it to be omitted if empty
	FirstName    *string            `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"required,email"` // email is a specific verification that checks for @ etc.
	Phone        *string            `json:"phone" validate:"required,min=2,max=10"`
	Token        *string            `json:"token,omitempty"`                                // "omitempty" skips this field if it's nil
	UserType     *string            `json:"user_type" validate:"required,oneof=ADMIN USER"` // "oneof" offers enum-like validation
	RefreshToken *string            `json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}

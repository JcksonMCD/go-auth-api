package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user document in MongoDB with validation tags
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FirstName    *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Password     *string            `bson:"password" json:"password" validate:"required,min=6"`
	Email        *string            `bson:"email" json:"email" validate:"required,email"` // email is a specific verification that checks for @ etc.
	Phone        *string            `bson:"phone" json:"phone" validate:"required,min=2,max=10"`
	Token        *string            `bson:"token,omitempty" json:"token,omitempty"`
	UserType     *string            `bson:"user_type" json:"user_type" validate:"required,oneof=ADMIN USER"` // "oneof" offers enum-like validation
	RefreshToken *string            `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	UserID       string             `bson:"user_id" json:"userI_id"`
}

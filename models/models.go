package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer Struct
type Customer struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	DateBirth string             `json:"dateBirth,omitempty"`
	Email     string             `json:"email,omitempty"`
	Phone     string             `json:"phone,omitempty"`
}

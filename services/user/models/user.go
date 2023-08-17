package models

import "time"

type User struct {
	ID          string      `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Email       string      `json:"email"`
	Password    string      `json:"-"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     Address     `json:"address"`
	Preferences Preferences `json:"preferences"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Address struct {
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	ZipCode   string  `json:"zipCode"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Preferences struct {
	Allergies        []string `json:"allergies"`
	DietaryChoice    string   `json:"dietaryChoice"`
	FavoriteCuisines []string `json:"favoriteCuisines"`
}

package models

// AvailableFund is the model for the response body when making the get available funds request
type AvailableFund struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

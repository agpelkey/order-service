package domain

type Entree struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Cost int `json:"cost"`
	Quantity int `json:"quantity"`
}

// model to create a new entree
type EntreeCreate struct {
	Name string `json:"name"`
	Cost int `json:"cost"`
}

// model to update entree
type EntreeUpdate struct {
	Name *string `json:"name"`
	Description *string `json:"description"`
	Cost *int `json:"cost"`
	Quantity *int `json:"quantity"`
}

// repository pattern for entree's
type EntreeService interface {
	// Create
	// GetByID
	// GetByName
	// Update
	// Delete
}



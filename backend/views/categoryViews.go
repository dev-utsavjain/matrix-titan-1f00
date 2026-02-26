package views

// CategoryResponse represents category response
type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt"`
}

package object

import "net/http"

// CategoryResponse represents response object
// of Category model
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Render processes object before rendered
func (res *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

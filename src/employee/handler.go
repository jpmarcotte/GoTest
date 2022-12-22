package employee

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler represents an adapter for the underlying data storage.
type Handler struct {
	store Store
}

// NewHandler returns a wrapper with the underlying storage injected.
func NewHandler() (*Handler, error) {
	// change which lines are commented to switch to a static in-memory store.
	//return &Handler{store: NewMemoryStore()}, nil
	store, err := NewSQLiteStore()
	return &Handler{store: store}, err
}

// GetAllEmployees outputs a JSON array of all employees currently in the store.
func (h *Handler) GetAllEmployees(ctx echo.Context) error {
	employees, err := h.store.GetAllEmployees()
	if err != nil {
		return ctx.JSONPretty(http.StatusInternalServerError, struct{
			Message string `json:"message"`
		}{Message:err.Error()}, "    ")
	}
	return ctx.JSONPretty(http.StatusOK, employees, "    ")
}

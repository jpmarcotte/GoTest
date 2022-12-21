package employee

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Handler struct {
	store Store
}

func NewHandler() (*Handler, error) {
	// Consider Store Repository if wanting to support more complex options
	//return &Handler{store: NewMemoryStore()}, nil
	store, err := NewSQLStore(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	return &Handler{store: store}, err
}

func (h *Handler) GetAllEmployees(ctx echo.Context) error {
	employees, err := h.store.GetAllEmployees()
	if err != nil {
		return ctx.JSONPretty(http.StatusInternalServerError, struct{
			Message string `json:"message"`
		}{Message:err.Error()}, "  ")
	}
	return ctx.JSONPretty(http.StatusOK, employees, "  ")
}
